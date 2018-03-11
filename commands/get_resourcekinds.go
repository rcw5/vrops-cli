package commands

import (
	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
)

func GetResourceKinds(adapterKind string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	resourceKinds, err := client.ResourceKinds(adapterKind)
	if err != nil {
		return err
	}

	presenter.PresentResourceKinds(resourceKinds)
	return nil
}
