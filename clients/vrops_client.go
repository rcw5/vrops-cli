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
	AdapterKinds() (models.AdapterKinds, error)
	ResourceKinds(string) ([]string, error)
	ResourcesForAdapterKind(string) (models.Resources, error)
	FindResource(string, string) (models.Resource, error)
	CreateStats(string, models.Stats) error
	GetStatsForResource(string, string, string) (models.ListStatsResponseValuesStatListStats, error)
	CreateResource(models.Resource) error
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

func (c VRopsClient) GetStatsForResource(adapterKind, resourceName, statKey string) (models.ListStatsResponseValuesStatListStats, error) {
	resource, err := c.FindResource(adapterKind, resourceName)
	if err != nil {
		return models.ListStatsResponseValuesStatListStats{}, err
	}
	request, err := http.NewRequest("GET", c.buildUrl(fmt.Sprintf("api/resources/%s/stats", resource.Identifier)), nil)
	if err != nil {
		return models.ListStatsResponseValuesStatListStats{}, err
	}

	response, err := c.do(request)
	if err != nil {
		return models.ListStatsResponseValuesStatListStats{}, err
	}

	data := models.ListStatsResponse{}
	if err := json.Unmarshal(response, &data); err != nil {
		return models.ListStatsResponseValuesStatListStats{}, fmt.Errorf("Cannot parse response: %s", err)
	}

	return data.Values[0].StatList.Stat, nil

}

func (c VRopsClient) FindResource(adapterKind, resourceName string) (models.Resource, error) {
	adapterKinds, err := c.AdapterKinds()
	if err != nil {
		return models.Resource{}, fmt.Errorf("Error retrieving adapterkinds: %s", err)
	}

	_, err = adapterKinds.FindAdapterKind(adapterKind)
	if err != nil {
		return models.Resource{}, err
	}

	resources, err := c.ResourcesForAdapterKind(adapterKind)
	if err != nil {
		return models.Resource{}, fmt.Errorf("Error retrieving resources: %s", err)
	}

	resource, err := resources.FindResource(resourceName)
	if err != nil {
		return models.Resource{}, err
	}
	return resource, nil
}

func (c VRopsClient) CreateResource(resource models.Resource) error {
	jsonEnc, err := json.Marshal(resource)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.buildUrl(fmt.Sprintf("api/resources/adapterkinds/%s", resource.ResourceKey.AdapterKindKey)), bytes.NewBuffer(jsonEnc))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	_, err = c.do(request)
	return err

}
func (c VRopsClient) CreateStats(resource string, stats models.Stats) error {
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

func (c VRopsClient) ResourcesForAdapterKind(adapterKind string) (models.Resources, error) {
	request, err := http.NewRequest("GET", c.buildUrl(fmt.Sprintf("api/adapterkinds/%s/resources", adapterKind)), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.do(request)
	if err != nil {
		return models.Resources{}, err
	}
	var data struct {
		PageInfo     models.PageInfo  `json:"pageInfo"`
		ResourceList models.Resources `json:"resourceList"`
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return models.Resources{}, err
	}
	if data.PageInfo.TotalCount > data.PageInfo.PageSize {
		return models.Resources{}, errors.New("No support for result pagination yet, mate")
	}
	return data.ResourceList, nil
}

func (c VRopsClient) AdapterKinds() (models.AdapterKinds, error) {
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
		return []string{}, fmt.Errorf("Cannot parse response: %s", err)
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
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Request failed: %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
