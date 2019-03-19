package commands

import (
	"encoding/json"

	"github.com/topflight-technology/vrops-cli/clients"
	"github.com/topflight-technology/vrops-cli/models"
	"github.com/spf13/cobra"
)

var createStatsCmd = &cobra.Command{
	Use:   "stats [adapterkind] [resource]",
	Short: "Create stats for a resource",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmdErr = CreateStats(args[0], args[1], cmd.Flag("statsjson").Value.String(), client)
	},
}

func init() {
	createStatsCmd.Flags().String("statsjson", "", "JSON-encoded stats to be uploaded")
	createStatsCmd.MarkFlagRequired("statsjson")
	createCmd.AddCommand(createStatsCmd)
}

func CreateStats(adapterKind, resourceName, statsJson string, client clients.VRopsClientIntf) error {
	resource, err := client.FindResource(adapterKind, resourceName)
	if err != nil {
		return err
	}
	stats := []models.Stat{}
	err = json.Unmarshal([]byte(statsJson), &stats)
	if err != nil {
		return err
	}
	err = client.CreateStats(resource.Identifier, stats)
	if err != nil {
		return err
	}
	return nil
}
