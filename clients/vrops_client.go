package clients

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rcw5/vrops-cli/models"
)

//go:generate counterfeiter -o ../fakes/FakeVRopsClient.go --fake-name FakeVRopsClient . VRopsClientIntf
type VRopsClientIntf interface {
	AdapterKinds() (models.AdapterKinds, error)
	ResourceKinds(string) ([]string, error)
}

type VRopsClient struct {
	url      string
	username string
	password string
}

func NewVROpsClient(url, username, password string) VRopsClient {
	return VRopsClient{
		url:      url,
		username: username,
		password: password,
	}
}

func (c VRopsClient) AdapterKinds() (models.AdapterKinds, error) {
	response, err := c.performGet("api/adapterkinds")
	if err != nil {
		return models.AdapterKinds{}, err
	}
	var dat models.AdapterKinds
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	return dat, nil
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
	response, err := client.Do(req)
	if err != nil {
		return nil, err
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
