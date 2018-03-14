package commands

import (
	"fmt"
	"os"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
	"github.com/spf13/cobra"
)

var getResourceCmd = &cobra.Command{
	Use:   "resources [adapterkind]",
	Short: "List resources for a given adapterkind",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		presenter := presenters.NewPresenter(cmd.Flag("output").Value.String())
		err := GetResources(args[0], client, presenter)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	getCmd.AddCommand(getResourceCmd)
}

func GetResources(adapterKind string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	resources, err := client.ResourcesForAdapterKind(adapterKind)
	if err != nil {
		return err
	}

	presenter.PresentResources(resources)
	return nil
}
