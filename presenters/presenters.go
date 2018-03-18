package presenters

import (
	"os"

	"github.com/rcw5/vrops-cli/models"
)

//go:generate counterfeiter -o ../fakes/FakePresenter.go --fake-name FakePresenter . Presenter
type Presenter interface {
	PresentAdapterKinds(models.AdapterKinds)
	PresentResourceKinds([]string)
	PresentResources(models.Resources)
	PresentStats(models.ListStatsResponseValuesStatListStats)
}

func NewPresenter(presenterType string) Presenter {
	if presenterType == "json" {
		return JSONPresenter{Buffer: os.Stdout}
	} else {
		return TablePresenter{Buffer: os.Stdout}
	}
}
