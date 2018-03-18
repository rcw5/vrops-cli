package commands

import (
	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/presenters"
	"github.com/spf13/cobra"
)

var getStatsCmd = &cobra.Command{
	Use:   "stats [adapterkind] [resource]",
	Short: "List stats for a given adapterkind and resource",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		presenter := presenters.NewPresenter(cmd.Flag("output").Value.String())
		cmdErr = GetStats(args[0], args[1], cmd.Flag("stat-key").Value.String(), client, presenter)
	},
}

func init() {
	getStatsCmd.Flags().String("stat-key", "", "Only show statistics for this key (default: all keys shown)")
	getCmd.AddCommand(getStatsCmd)
}

func GetStats(adapterKind, resource, statKey string, client clients.VRopsClientIntf, presenter presenters.Presenter) error {
	stats, err := client.GetStatsForResource(adapterKind, resource, statKey)
	if err != nil {
		return err
	}
	presenter.PresentStats(stats)

	return nil
}
