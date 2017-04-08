package vebra

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"vebra/vebra/model"
)

var host = `http://webservices.vebra.com`

type DataGetter func(api *API, URL string, OUT interface{}) (err error)
type TokenStorage interface {
	Save(token string) error
	Load() (string, error)
}

type API struct {
	baseUrl      string
	dataFeedId   string
	username     string
	password     string
	tokenStorage TokenStorage
	feedVersion  int
	client       *http.Client
	GetDataFunc  DataGetter
}

func Create(dataFeedId string, username string, password string, tokenStorage TokenStorage) (api *API) {
	api = &API{}
	api.feedVersion = 10
	api.dataFeedId = dataFeedId
	api.username = username
	api.password = password
	api.tokenStorage = tokenStorage

	api.client = &http.Client{}
	api.baseUrl = fmt.Sprintf("export/%s/v%d", api.dataFeedId, api.feedVersion)
	return api
}

func (api API) execute(url string, secondAttempt bool) (res *http.Response, err error) {
	if url == "" {
		return nil, fmt.Errorf("URL must be a non-empty string")
	}

	// START
	callUrl := fmt.Sprintf("%s/%s/%s", host, api.baseUrl, url)

	if res, err = api.tokenAuth(callUrl); err != nil {
		return res, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res, nil
	case http.StatusUnauthorized:
		if secondAttempt {
			return nil, fmt.Errorf("Unauthorized")
		}

		// try again with basic auth
		if res, err = api.basicAuth(callUrl); err != nil {
			return nil, err
		} else if res.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("Unauthorized - error authenticating. An active token hasn't expired or invalid credentials have been provided")
		}
		return res, err
	default:
		return nil, fmt.Errorf("ERR: %s", http.StatusText(res.StatusCode))
	}
}

func (api *API) tokenAuth(callUrl string) (res *http.Response, err error) {
	loadedToken, err := api.tokenStorage.Load()

	if err != nil {
		return api.basicAuth(callUrl)
	}

	req, _ := http.NewRequest("GET", callUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", string(loadedToken)))
	res, err = api.client.Do(req)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (api *API) basicAuth(callUrl string) (res *http.Response, err error) {
	req, _ := http.NewRequest("GET", callUrl, nil)

	req.SetBasicAuth(api.username, api.password)
	res, err = api.client.Do(req)

	if err != nil {
		return res, err
	}

	// save the token
	if token := res.Header.Get("Token"); token != "" {
		err = api.tokenStorage.Save(token)
	}
	return res, nil
}

// GetBranches returns an array of branches
func (api *API) GetBranches() (*model.BranchSummaries, error) {
	branches := &model.BranchSummaries{}

	if err := api.GetDataFunc(api, "branch", branches); err != nil {
		return nil, err
	}

	return branches, nil
}

// GetBranches returns an array of branches
func (api *API) GetBranch(branchSummary model.BranchSummary) (*model.Branch, error) {

	ClientID, err := branchSummary.GetClientID()

	if err != nil {
		return nil, err
	}

	branch := &model.Branch{}
	branch.BranchID = branchSummary.BranchID
	branch.FirmID = branchSummary.FirmID
	branch.ClientID = ClientID

	if err := api.GetDataFunc(api, fmt.Sprintf("branch/%d", ClientID), branch); err != nil {
		return nil, err
	}

	return branch, err
}

func (api *API) GetPropertyList(ClientID int) (properties *model.PropertySummaries, err error) {
	properties = &model.PropertySummaries{}

	if err := api.GetDataFunc(api, fmt.Sprintf("branch/%d/property", ClientID), properties); err != nil {
		return nil, err
	}

	return properties, err
}

func (api *API) GetProperty(clientId int, propertyId int) (property *model.Property, err error) {
	property = &model.Property{}
	if err = api.GetDataFunc(api, fmt.Sprintf(`branch/%d/property/%d`, clientId, propertyId), property); err != nil {
		return nil, err
	}

	return property, err
}

func (api *API) GetChangedProperties(since time.Time) (properties *[]model.ChangedPropertySummary, err error) {
	propsums := &model.ChangedPropertySummaries{}
	if err = api.GetDataFunc(api, fmt.Sprintf(`property/%s`, since.Format(`2006/01/02/15/04/05`)), propsums); err != nil {
		return nil, err
	}

	return &propsums.PropertySummaries, err
}

func (api *API) GetChangedFiles(since time.Time) (files *[]model.ChangedFileSummary, err error) {
	filesums := &model.ChangedFilesSummaries{}
	if err = api.GetDataFunc(api, fmt.Sprintf(`files/%s`, since.Format(`2006/01/02/15/04/05`)), filesums); err != nil {
		return nil, err
	}

	return &filesums.Files, err
}

func NewRemoteFileGetter() DataGetter {
	return func(api *API, URL string, OUT interface{}) (err error) {
		var response *http.Response

		if response, err = api.execute(URL, false); err != nil {
			return err
		}

		defer response.Body.Close()
		bodybuffer := new(bytes.Buffer)

		if _, err = bodybuffer.ReadFrom(response.Body); err != nil {
			return err
		}

		err = xml.Unmarshal(bodybuffer.Bytes(), OUT)
		return err
	}
}

func NewRemoteFileGetterLocalWriter(outdir string) DataGetter {
	if _, err := os.Stat(outdir); err != nil {
		log.Fatalf("ERR: Output directory [%s] does not exist.")
	}
	return func(api *API, URL string, OUT interface{}) (err error) {
		fmt.Println("Got URL: %s", URL)
		putpath := outdir
		outfile := fmt.Sprintf("%s%s", path.Join(putpath, URL), ".xml")
		var response *http.Response

		if response, err = api.execute(URL, false); err != nil {
			return err
		}

		defer response.Body.Close()
		bodybuffer := new(bytes.Buffer)

		ind := strings.LastIndex(URL, "/")

		log.Printf("Got ind [%d] in URL [%s]", ind, URL)
		if ind > -1 {
			putpath = path.Join(outdir, URL[:ind])
			outfile = fmt.Sprintf("%s%s", path.Join(putpath, URL[ind+1:]), ".xml")
			log.Printf("Creating directory: [%s]", putpath)
			if err := os.MkdirAll(putpath, 0777); err != nil {
				log.Fatalf("ERR: Error creating directory [%s]: %s", putpath, err)
			}

		}

		if _, err = bodybuffer.ReadFrom(response.Body); err != nil {
			log.Printf("%s", err)
			return err
		}

		log.Printf("Writing to file: [%s]", outfile)
		if err := ioutil.WriteFile(outfile, bodybuffer.Bytes(), 0777); err != nil {
			log.Fatalf("ERR: Error creating file [%s]: %s", outfile, err)
		}

		err = xml.Unmarshal(bodybuffer.Bytes(), OUT)
		return err
	}
}

func NewLocalFileGetter(outdir string) DataGetter {
	if _, err := os.Stat(outdir); err != nil {
		log.Fatalf("ERR: Data directory [%s] does not exist.")
	}
	return func(api *API, URL string, OUT interface{}) (err error) {
		inpath := outdir
		in := fmt.Sprintf("%s%s", path.Join(inpath, URL), ".xml")
		ind := strings.LastIndex(URL, "/")

		if ind > -1 {
			inpath = path.Join(outdir, URL[:ind])
			in = fmt.Sprintf("%s%s", path.Join(inpath, URL[ind+1:]), ".xml")
		}

		if _, err := os.Stat(in); err != nil {
			return err
		}

		file, err := ioutil.ReadFile(in)

		err = xml.Unmarshal(file, OUT)
		return err
	}
}
