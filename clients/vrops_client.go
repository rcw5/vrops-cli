package clients

import (
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

func (c VRopsClient) ResourcesForAdapterKind(adapterKind string) ([]models.Resource, error) {
	response, err := c.performGet(fmt.Sprintf("api/adapterkinds/%s/resources", adapterKind))
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
	if data.PageInfo.TotalCount != 1 {
		return []models.Resource{}, errors.New("No support for result pagination yet, mate")
	}
	return data.ResourceList, nil
}

func (c VRopsClient) AdapterKinds() ([]models.AdapterKind, error) {
	response, err := c.performGet("api/adapterkinds")
	if err != nil {
		return []models.AdapterKind{}, err
	}

	var data struct {
		PageInfo     models.PageInfo      `json:"pageInfo"`
		AdapterKinds []models.AdapterKind `json:"adapter-kind"`
	}

	if err := json.Unmarshal(response, &data); err != nil {
		return []models.AdapterKind{}, err
	}
	if data.PageInfo.TotalCount != 1 {
		return []models.AdapterKind{}, errors.New("No support for result pagination yet, mate")
	}
	return data.AdapterKinds, nil
}

func (c VRopsClient) ResourceKinds(adapterKind string) ([]string, error) {
	response, err := c.performGet(fmt.Sprintf("api/adapterkinds/%s", adapterKind))
	if err != nil {
		return []string{}, err
	}
	var dat models.AdapterKind
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	return dat.ResourceKinds, nil
}

func (c VRopsClient) performGet(uri string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/suite-api/%s", c.url, uri), nil)
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)
	if c.verbose {
		requestOutput, _ := httputil.DumpRequest(req, false)
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
