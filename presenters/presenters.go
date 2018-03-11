package presenters

import (
	"os"

	"github.com/rcw5/vrops-cli/models"
)

//go:generate counterfeiter -o ../fakes/FakePresenter.go --fake-name FakePresenter . Presenter
type Presenter interface {
	PresentAdapterKinds([]models.AdapterKind)
	PresentResourceKinds([]string)
	PresentResources([]models.Resource)
}

func NewPresenter(presenterType string) Presenter {
	if presenterType == "json" {
		return JSONPresenter{Buffer: os.Stdout}
	} else {
		return TablePresenter{Buffer: os.Stdout}
	}
}
