package models

type AdapterKind struct {
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	AdapterKindType string   `json:"adapterKindType"`
	DescribeVersion int      `json:"describeVersion"`
	ResourceKinds   []string `json:"resourceKinds"`
}

type AdapterKinds struct {
	AdapterKind []AdapterKind `json:"adapter-kind"`
}
