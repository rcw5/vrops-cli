package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/rcw5/vrops-cli/models"
)

//go:generate counterfeiter -o ../fakes/FakeVRopsClient.go --fake-name FakeVRopsClient . VRopsClientIntf
type VRopsClientIntf interface {
	AdapterKinds() ([]models.AdapterKind, error)
	ResourceKinds(string) ([]string, error)
	ResourcesForAdapterKind(string) ([]models.Resource, error)
	CreateStats(string, []models.Stat) error
}

type VRopsClient struct {
	url      string
	username string
	password string
	verbose  bool
}

func NewVROpsClient(url, username, password string, verbose bool) VRopsClient {
	return VRopsClient{
		url:      url,
		username: username,
		password: password,
		verbose:  verbose,
	}
}

func (c VRopsClient) CreateStats(resource string, stats []models.Stat) error {
	data := struct {
		Stats []models.Stat `json:"stat-content"`
	}{
		Stats: stats,
	}

	jsonEnc, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.buildUrl(fmt.Sprintf("api/resources/%s/stats", resource)), bytes.NewBuffer(jsonEnc))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	_, err = c.do(request)
	return err
}

func (c VRopsClient) ResourcesForAdapterKind(adapterKind string) ([]models.Resource, error) {
	request, err := http.NewRequest("GET", c.buildUrl(fmt.Sprintf("api/adapterkinds/%s/resources", adapterKind)), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.do(request)
	if err != nil {
		return []models.Resource{}, err
	}
	var data struct {
		PageInfo     models.PageInfo   `json:"pageInfo"`
		ResourceList []models.Resource `json:"resourceList"`
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return []models.Resource{}, err
	}
	if data.PageInfo.TotalCount > 1 {
		return []models.Resource{}, errors.New("No support for result pagination yet, mate")
	}
	return data.ResourceList, nil
}

func (c VRopsClient) AdapterKinds() ([]models.AdapterKind, error) {
	request, err := http.NewRequest("GET", c.buildUrl("api/adapterkinds"), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.do(request)
	if err != nil {
		return []models.AdapterKind{}, err
	}

	var data struct {
		AdapterKinds []models.AdapterKind `json:"adapter-kind"`
	}

	if err := json.Unmarshal(response, &data); err != nil {
		return []models.AdapterKind{}, err
	}
	return data.AdapterKinds, nil
}

func (c VRopsClient) ResourceKinds(adapterKind string) ([]string, error) {
	request, err := http.NewRequest("GET", c.buildUrl(fmt.Sprintf("api/adapterkinds/%s", adapterKind)), nil)
	if err != nil {
		return []string{}, err
	}

	response, err := c.do(request)
	if err != nil {
		return []string{}, err
	}
	var dat models.AdapterKind
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	return dat.ResourceKinds, nil
}

func (c VRopsClient) buildUrl(uri string) string {
	return fmt.Sprintf("%s/suite-api/%s", c.url, uri)
}

func (c VRopsClient) do(req *http.Request) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)

	if c.verbose {
		dumpBody := req.Method != "GET"
		requestOutput, _ := httputil.DumpRequest(req, dumpBody)
		fmt.Println(string(requestOutput))
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.verbose {
		responseOutput, _ := httputil.DumpResponse(response, true)
		fmt.Println(string(responseOutput))

	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Request failed: %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
