package commands

import (
	"github.com/topflight-technology/vrops-cli/presenters"
	"github.com/spf13/cobra"

	"github.com/topflight-technology/vrops-cli/clients"
)

var getAdapterKindsCmd = &cobra.Command{
	Use:   "adapterkinds",
	Short: "Get all adapterkinds",
	Run: func(cmd *cobra.Command, args []string) {
		presenter := presenters.NewPresenter(cmd.Flag("output").Value.String())
		cmdErr = GetAdapterKinds(client, presenter)
	},
}

func init() {
	getCmd.AddCommand(getAdapterKindsCmd)
}

func GetAdapterKinds(client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	adapterKinds, err := client.AdapterKinds()
	if err != nil {
		return err
	}
	presenter.PresentAdapterKinds(adapterKinds)
	return nil
}
