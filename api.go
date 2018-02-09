package main

import (
	"net/http"
	"time"
	"fmt"
	"io"
	"bytes"
	"encoding/xml"
	"strconv"
)

const (
	HeaderAuthorizationKey      = "Authorization"
	HeaderIfModifiedSinceKey    = "If-Modified-Since"
	HeaderTokenKey              = "Token"
	HeaderTokenAuthFormatString = "Basic %s"
)

type TokenStorage interface {
	Save(token token) error
	Load() (*token, error)
}

type Api struct {
	dataFeedId   string
	credentials  *credentials
	tokenStorage TokenStorage
	Error        error
	StatusCode   int
}

func NewApi(dataFeedId string, username string, password string) (*Api) {
	return &Api{
		dataFeedId: dataFeedId,
		credentials: &credentials{
			userName: username,
			password: password,
		},
	}
}

func (api *Api) SetTokenStorage(tokenStorage TokenStorage) {
	api.tokenStorage = tokenStorage
}

// GetBranches returns an array of branches
func (api *Api) GetBranches() (*BranchSummaries, error) {
	branches := new(BranchSummaries)
	branchesURLBuilder := new(URLGetBranchesBuilder)
	branchesURLBuilder.SetDataFeedID(api.dataFeedId)
	if err := api.doRequest(branchesURLBuilder, branches); err != nil {
		return nil, err
	}
	return branches, nil
}

func (api *Api) GetBranch(branchSummary *BranchSummary) (branch *Branch, err error) {
	branch = new(Branch)
	branchURLBuilder := new(URLGetBranchBuilder)
	branchURLBuilder.SetDataFeedID(api.dataFeedId)
	branchURLBuilder.SetClientID(branchSummary.GetClientIDString())
	if err := api.doRequest(branchURLBuilder, branch); err != nil {
		return nil, err
	}
	return branch, nil
}

func (api *Api) GetProperties(branchSummary *BranchSummary) (properties *PropertySummaries, err error) {
	properties = new(PropertySummaries)
	propertiesURLBuilder := new(URLGetPropertiesBuilder)
	propertiesURLBuilder.SetDataFeedID(api.dataFeedId)
	propertiesURLBuilder.SetClientID(branchSummary.GetClientIDString())
	if err := api.doRequest(propertiesURLBuilder, properties); err != nil {
		return nil, err
	}
	return properties, nil
}

func (api *Api) GetProperty(branchSummary *BranchSummary, summary PropertySummary) (property *Property, err error) {
	property = new(Property)
	propertyURLBuilder := new(URLGetPropertyBuilder)
	propertyURLBuilder.SetDataFeedID(api.dataFeedId)
	propertyURLBuilder.SetClientID(branchSummary.GetClientIDString())
	propertyURLBuilder.SetPropertyID(strconv.Itoa(int(summary.PropertyID)))
	if err := api.doRequest(propertyURLBuilder, property); err != nil {
		return nil, err
	}
	return property, nil
}


func (api *Api) GetPropertyFromChangedFileSummary(summary ChangedFileSummary) (property *Property, err error) {
	property = new(Property)
	propertyURLBuilder := new(ChangedPropertyURLBuilder)
	propertyURLBuilder.SetURL(summary.PropUrl)
	if err := api.doRequest(propertyURLBuilder, property); err != nil {
		return nil, err
	}
	return property, nil
}

func (api *Api) GetChangedProperties(since time.Time) (properties *ChangedPropertySummaries, err error) {
	properties = new(ChangedPropertySummaries)
	propertiesURLBuilder := new(URLGetChangedPropertiesBuilder)
	propertiesURLBuilder.SetDataFeedID(api.dataFeedId)
	propertiesURLBuilder.SetSince(since)
	if err = api.doRequest(propertiesURLBuilder, properties); err != nil {
		return nil, err
	}
	return properties, nil
}

func (api *Api) GetChangedProperty(changedProperty *ChangedPropertySummary) (property *Property, err error) {
	if changedProperty.LastAction == Deleted {
		return nil, fmt.Errorf("property [%s] has been deleted", strconv.Itoa(int(changedProperty.PropertyID)))
	}
	property = new(Property)
	propertiesURLBuilder := new(ChangedPropertyURLBuilder)
	propertiesURLBuilder.SetURL(changedProperty.Url)
	if err = api.doRequest(propertiesURLBuilder, property); err != nil {
		return nil, err
	}
	return property, nil
}

func (api *Api) GetChangedFiles(since time.Time) (changedFiles *ChangedFilesSummaries, err error) {
	changedFiles = new(ChangedFilesSummaries)
	urlGetChangedFilesBuilder := new(URLGetChangedFilesBuilder)
	urlGetChangedFilesBuilder.SetDataFeedID(api.dataFeedId)
	urlGetChangedFilesBuilder.SetSince(since)
	if err = api.doRequest(urlGetChangedFilesBuilder, changedFiles); err != nil {
		return nil, err
	}
	return changedFiles, nil
}

func (api *Api) doRequest(urlBuilder URLBuilder, out interface{}) (err error) {
	requestor := buildRequestor(api.dataFeedId, api.credentials)
	if api.tokenStorage != nil {
		if requestor.token, err = api.tokenStorage.Load(); err != nil {
			return err
		}
	}
	requestor.urlBuilder = urlBuilder
	for ; requestor.attempts < 2; requestor.attempts++ {
		requestor.buildRequest()
		requestor.doRequest()
		requestor.saveTokenIfExists(api.tokenStorage)
		requestor.handleErrors()
		if requestor.response.StatusCode == http.StatusOK {
			return requestor.unmarshal(out)
		}
	}
	api.StatusCode = requestor.response.StatusCode
	api.Error = fmt.Errorf(requestor.response.Status)
	return api.Error
}

func (api *Api) doRequestSince(urlBuilder URLBuilder, out interface{}, since *time.Time) (err error) {
	requestor := buildRequestor(api.dataFeedId, api.credentials)
	requestor.setIfModifiedSince(since)
	if api.tokenStorage != nil {
		if requestor.token, err = api.tokenStorage.Load(); err != nil {
			return err
		}
	}
	requestor.urlBuilder = urlBuilder
	for ; requestor.attempts < 2; requestor.attempts++ {
		requestor.buildRequest()
		requestor.doRequest()
		requestor.saveTokenIfExists(api.tokenStorage)
		requestor.handleErrors()
		api.StatusCode = requestor.response.StatusCode
		if requestor.response.StatusCode == http.StatusOK {
			return requestor.unmarshal(out)
		}
	}
	api.Error = fmt.Errorf(requestor.response.Status)
	return api.Error
}

type credentials struct {
	userName string
	password string
}

type token struct {
	tokenString string
	isEmpty     bool
	isValid     bool
	timeSet     time.Time
}

func Token(tokenString string) *token {
	token := new(token)
	token.tokenString = tokenString
	token.isValid = true
	token.timeSet = time.Now()
	return token
}

func (token *token) IsValid() bool {
	return token.isValid
}

func (token *token) Invalidate() {
	token.isValid = false
	token.tokenString = ""
}

type requestor struct {
	dataFeedID  string
	credentials *credentials
	token       *token
	urlBuilder  URLBuilder
	header      http.Header
	since       *time.Time
	request     *http.Request
	response    *http.Response
	body        io.Reader
	err         error
	attempts    int
}

func buildRequestor(dataFeedId string, credentials *credentials) (*requestor) {
	return &requestor{
		dataFeedID:  dataFeedId,
		credentials: credentials,
		token:       &token{},
		urlBuilder:  nil,
		header:      http.Header{},
	}
}

func (requestor *requestor) doRequest() {
	requestor.response, requestor.err = (&http.Client{}).Do(requestor.request)
	if requestor.err != nil {
		panic(requestor.err)
	}
}

func (requestor *requestor) saveTokenIfExists(tokenStorage TokenStorage) {
	if token := requestor.response.Header.Get(HeaderTokenKey); token != "" {
		requestor.token = Token(token)
		if tokenStorage != nil {
			tokenStorage.Save(*requestor.token)
		}
	}
}
func (requestor *requestor) handleErrors() {
	switch requestor.response.StatusCode {
	case http.StatusUnauthorized:
		requestor.token.Invalidate()
	default:
	}
}

func (requestor *requestor) setIfModifiedSince(since *time.Time) {
	if since != nil {
		requestor.header.Add(HeaderIfModifiedSinceKey, since.Format(time.RFC1123))
	}
}

func (requestor *requestor) setHeaderAttribute(key string, value string) {
	requestor.header.Add(key, value)
}

func (requestor *requestor) setBasicAuth() {
	requestor.request.SetBasicAuth(requestor.credentials.userName, requestor.credentials.password)
}

func (requestor *requestor) setAuthenticationToken() {
	requestor.request.Header.Add(HeaderAuthorizationKey, fmt.Sprintf(HeaderTokenAuthFormatString, requestor.token.tokenString))
}

func (requestor *requestor) buildRequest() (err error) {
	requestor.request, err = http.NewRequest(http.MethodGet, requestor.urlBuilder.Build(), requestor.body)
	for key, val := range requestor.header {
		requestor.request.Header[key] = val
	}
	requestor.setAuthenticationMethod()
	return err
}

func (requestor *requestor) setAuthenticationMethod() {
	if requestor.token.IsValid() {
		requestor.setAuthenticationToken()
		return
	}
	requestor.setBasicAuth()
}

func (requestor *requestor) unmarshal(out interface{}) error {
	defer requestor.response.Body.Close()
	bodyBuffer := new(bytes.Buffer)
	if _, err := bodyBuffer.ReadFrom(requestor.response.Body); err != nil {
		return err
	}
	err := xml.Unmarshal(bodyBuffer.Bytes(), out)
	return err
}
