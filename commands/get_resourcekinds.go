package commands

import (
	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
	"github.com/spf13/cobra"
)

var getResourceKindsCmd = &cobra.Command{
	Use:   "resourcekinds [adapterkind]",
	Short: "List resourceskinds for a given adapterkind",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		presenter := presenters.NewPresenter(cmd.Flag("output").Value.String())
		cmdErr = GetResourceKinds(args[0], client, presenter)
	},
}

func init() {
	getCmd.AddCommand(getResourceKindsCmd)
}

func GetResourceKinds(adapterKind string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	resourceKinds, err := client.ResourceKinds(adapterKind)
	if err != nil {
		return err
	}

	presenter.PresentResourceKinds(resourceKinds)
	return nil
}
