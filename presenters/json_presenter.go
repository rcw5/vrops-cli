package presenters

import (
	"encoding/json"
	"io"

	"github.com/rcw5/vrops-cli/models"
)

type JSONPresenter struct {
	Buffer io.Writer
}

func (j JSONPresenter) PresentAdapterKinds(adapterKinds []models.AdapterKind) {
	j.jsonify(adapterKinds)
}

func (j JSONPresenter) PresentResourceKinds(resourceKinds []string) {
	j.jsonify(resourceKinds)
}

func (j JSONPresenter) PresentResources(resources []models.Resource) {
	j.jsonify(resources)
}

func (j JSONPresenter) jsonify(intf interface{}) {
	output, _ := json.Marshal(intf)
	j.Buffer.Write(output)
}
