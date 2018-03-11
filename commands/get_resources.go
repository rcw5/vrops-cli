package commands

import (
	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
)

func GetResources(adapterKind string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	resources, err := client.ResourcesForAdapterKind(adapterKind)
	if err != nil {
		return err
	}

	presenter.PresentResources(resources)
	return nil
}
