package model

import (
	"net/http"
	"time"
	"fmt"
	"io"
	"bytes"
	"encoding/xml"
)

type TokenStorage interface {
	Save(token token) error
	Load() (token, error)
}

type api struct {
	dataFeedId  string
	credentials *credentials
	tokenStorage TokenStorage
}

func Api(dataFeedId string, username string, password string) (api *api) {
	return &api{
		dataFeedId: dataFeedId,
		credentials: &credentials{
			userName: username,
			password: password,
		},
	}
}

func (api *api) SetTokenStorage(tokenStorage TokenStorage) {
	api.tokenStorage = tokenStorage
}

// GetBranches returns an array of branches
func (api *api) GetBranches() (*BranchSummaries, error) {
	branches := new(BranchSummaries)
	branchesURLBuilder := new(URLGetBranchesBuilder)
	branchesURLBuilder.SetDataFeedID(api.dataFeedId)
	if err := api.doRequest(branchesURLBuilder, branches); err != nil {
		return nil, err
	}
	return branches, nil
}

func (api *api) GetBranch(branchSummary *BranchSummary) (branch *Branch, err error) {
	var clientID int
	if clientID, err = branchSummary.GetClientID(); err != nil {
		return nil, err
	}
	branchURLBuilder := new(URLGetBranchBuilder)
	branchURLBuilder.SetDataFeedID(api.dataFeedId)
	branchURLBuilder.SetClientID(string(clientID))
	if err := api.doRequest(branchURLBuilder, branch); err != nil {
		return nil, err
	}
	return branch, nil
}

func (api *api) GetProperties(branchSummary *BranchSummary) (properties *PropertySummaries, err error) {
	var clientID int
	if clientID, err = branchSummary.GetClientID(); err != nil {
		return nil, err
	}
	propertiesURLBuilder := new(URLGetPropertiesBuilder)
	propertiesURLBuilder.SetDataFeedID(api.dataFeedId)
	propertiesURLBuilder.SetClientID(string(clientID))
	if err := api.doRequest(propertiesURLBuilder, properties); err != nil {
		return nil, err
	}
	return properties, nil
}

func (api *api) GetProperty(branchSummary *BranchSummary, summary PropertySummary) (property *Property, err error) {
	var clientID int
	if clientID, err = branchSummary.GetClientID(); err != nil {
		return nil, err
	}
	propertyURLBuilder := new(URLGetPropertyBuilder)
	propertyURLBuilder.SetDataFeedID(api.dataFeedId)
	propertyURLBuilder.SetClientID(string(clientID))
	propertyURLBuilder.SetPropertyID(string(summary.PropertyID))
	if err := api.doRequest(propertyURLBuilder, property); err != nil {
		return nil, err
	}
	return property, nil
}

func (api *api) GetChangedProperties(since time.Time) (properties *PropertySummaries, err error) {
	propertiesURLBuilder := new(URLGetChangedPropertiesBuilder)
	propertiesURLBuilder.SetDataFeedID(api.dataFeedId)
	propertiesURLBuilder.SetSince(since)
	if err = api.doRequest(propertiesURLBuilder, properties); err != nil {
		return nil, err
	}
	return properties, nil
}

func (api *api) GetChangedFiles(since time.Time) (changedFiles *ChangedFilesSummaries, err error) {
	urlGetChangedFilesBuilder := new(URLGetChangedFilesBuilder)
	urlGetChangedFilesBuilder.SetDataFeedID(api.dataFeedId)
	urlGetChangedFilesBuilder.SetSince(since)
	if err = api.doRequest(urlGetChangedFilesBuilder, changedFiles); err != nil {
		return nil, err
	}
	return changedFiles, nil
}

func (api *api)doRequest(urlBuilder URLBuilder, out interface{}) error {
	requestor := buildRequestor(api.dataFeedId, api.credentials)
	requestor.urlBuilder = urlBuilder
	for ;requestor.attempts < 2;requestor.attempts++ {
		requestor.buildRequest()
		requestor.setAuthenticationMethod()
		requestor.doRequest()
		requestor.saveTokenIfExists(api.tokenStorage)
		requestor.handleErrors()
		if requestor.response.StatusCode == http.StatusOK {
			return requestor.unmarshal(out)
		}
	}
	return fmt.Errorf(requestor.response.Status)
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
	requestor.response, requestor.err = http.DefaultClient.Do(requestor.request)
}

func (requestor *requestor) setIfModifiedSince(since *time.Time) {
	if since != nil {
		requestor.header.Add("If-Modified-Since", since.Format(time.RFC1123))
	}
}

func (requestor *requestor) setHeaderAttribute(key string, value string) {
	requestor.header.Add(key, value)
}

func (requestor *requestor) setBasicAuth() {
	requestor.request.SetBasicAuth(requestor.credentials.userName, requestor.credentials.password)
}

func (requestor *requestor) setAuthenticationToken() {
	requestor.header.Add("Authorization", fmt.Sprintf("Basic %s", string(requestor.token.tokenString)))
}

func (requestor *requestor) saveTokenIfExists(tokenStorage TokenStorage) {
	if token := requestor.response.Header.Get("Token"); token != "" {
		requestor.token = Token(token)
		if tokenStorage != nil {
			tokenStorage.Save(*requestor.token)
		}
	}
}

func (requestor *requestor) buildRequest() (err error) {
	requestor.request, err = http.NewRequest(http.MethodGet, requestor.urlBuilder.Build(), requestor.body)
	requestor.request.Header = requestor.header
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

func (requestor *requestor) handleErrors() {
	switch requestor.response.StatusCode {
	case http.StatusUnauthorized:
		requestor.token.Invalidate()
	default:
	}
}
