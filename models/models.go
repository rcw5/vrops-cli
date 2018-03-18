package models

import "fmt"

type PageInfo struct {
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}

type AdapterKinds []AdapterKind

func (r AdapterKinds) FindAdapterKind(name string) (AdapterKind, error) {
	for _, v := range r {
		if v.Name == name {
			return v, nil
		}
	}
	return AdapterKind{}, fmt.Errorf("Cannot find adapterkind: %s", name)
}

type AdapterKind struct {
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	AdapterKindType string   `json:"adapterKindType"`
	DescribeVersion int      `json:"describeVersion"`
	ResourceKinds   []string `json:"resourceKinds"`
}

type ResourceIdentifier struct {
	IdentifierType ResourceIdentifierType `json:"identifierType"`
	Value          string                 `json:"value"`
}

type ResourceKey struct {
	Name                string               `json:"name"`
	AdapterKindKey      string               `json:"adapterKindKey"`
	ResourceKindKey     string               `json:"resourceKindKey"`
	ResourceIdentifiers []ResourceIdentifier `json:"resourceIdentifiers"`
}

type ResourceIdentifierType struct {
	Name               string `json:"name"`
	DataType           string `json:"dataType"`
	IsPartOfUniqueness bool   `json:"isPartOfUniqueness"`
}

type ResourceStatusState struct {
	AdapterInstanceID string `json:"adapterInstanceId"`
	ResourceStatus    string `json:"resourceStatus"`
	ResourceState     string `json:"resourceState"`
	StatusMessage     string `json:"statusMessage"`
}
type ResourceBadge struct {
	Type  string  `json:"type"`
	Color string  `json:"color"`
	Score float32 `json:"score"`
}

type Resource struct {
	Description          string                `json:"description"`
	CreationTime         int64                 `json:"creationTime,omitempty"`
	ResourceKey          ResourceKey           `json:"resourceKey"`
	ResourceStatusStates []ResourceStatusState `json:"resourceStatusStates"`
	ResourceHealth       string                `json:"resourceHealth,omitempty"`
	ResourceHealthValue  float32               `json:"resourceHealthValue,omitempty"`
	DtEnabled            bool                  `json:"dtEnabled"`
	MonitoringInterval   int                   `json:"monitoringInterval"`
	Badges               []ResourceBadge       `json:"badges"`
	Identifier           string                `json:"identifier"`
}

type Resources []Resource

func (r Resources) FindResource(name string) (Resource, error) {
	for _, v := range r {
		if v.ResourceKey.Name == name {
			return v, nil
		}
	}
	return Resource{}, fmt.Errorf("Cannot find resource: %s", name)
}

type IntervalUnit struct {
	Quantifier int `json:"quantifier"`
}

type Stat struct {
	StatKey    string    `json:"statKey"`
	Timestamps []int64   `json:"timestamps"`
	Data       []float64 `json:"data,omitempty"`
}

type Stats []Stat

type ListStatsResponse struct {
	Values []ListStatsResponseValues `json:"values"`
}

type ListStatsResponseValues struct {
	ResourceID string                          `json:"resourceId"`
	StatList   ListStatsResponseValuesStatList `json:"stat-list"`
}

type ListStatsResponseValuesStatList struct {
	Stat ListStatsResponseValuesStatListStats `json:"stat"`
}

type ListStatsResponseValuesStatListStat struct {
	StatKey      ListStatsResponseValuesStatListStatStatKey `json:"statKey"`
	Timestamps   []int64                                    `json:"timestamps"`
	Data         []float64                                  `json:"data,omitempty"`
	IntervalUnit IntervalUnit                               `json:"intervalUnit,omitempty"`
}
type ListStatsResponseValuesStatListStatStatKey struct {
	Key string `json:"key"`
}

type ListStatsResponseValuesStatListStats []ListStatsResponseValuesStatListStat
