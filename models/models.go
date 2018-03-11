package models

type PageInfo struct {
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
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
	CreationTime         int64                 `json:"creationTime"`
	ResourceKey          ResourceKey           `json:"resourceKey"`
	ResourceStatusStates []ResourceStatusState `json:"resourceStatusStates"`
	ResourceHealth       string                `json:"resourceHealth"`
	ResourceHealthValue  float32               `json:"resourceHealthValue"`
	DtEnabled            bool                  `json:"dtEnabled"`
	MonitoringInterval   int                   `json:"monitoringInterval"`
	Badges               []ResourceBadge       `json:"badges"`
	Identifier           string                `json:"identifier"`
}