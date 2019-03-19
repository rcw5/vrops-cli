package presenters

import (
	"encoding/json"
	"io"

	"github.com/topflight-technology/vrops-cli/models"
)

type JSONPresenter struct {
	Buffer io.Writer
}

func (j JSONPresenter) PresentAdapterKinds(adapterKinds models.AdapterKinds) {
	j.jsonify(adapterKinds)
}

func (j JSONPresenter) PresentResourceKinds(resourceKinds []string) {
	j.jsonify(resourceKinds)
}

func (j JSONPresenter) PresentResources(resources models.Resources) {
	j.jsonify(resources)
}

func (j JSONPresenter) PresentStats(stats models.ListStatsResponseValuesStatListStats) {
	j.jsonify(stats)
}

func (j JSONPresenter) jsonify(intf interface{}) {
	output, _ := json.Marshal(intf)
	j.Buffer.Write(output)
}
