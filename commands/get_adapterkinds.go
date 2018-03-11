package commands

import (
	"github.com/rcw5/vrops-cli/presenters"

	"github.com/rcw5/vrops-cli/clients"
)

func GetAdapterKinds(client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	adapterKinds, err := client.AdapterKinds()
	if err != nil {
		return err
	}
	presenter.PresentAdapterKinds(adapterKinds)
	return nil
}
