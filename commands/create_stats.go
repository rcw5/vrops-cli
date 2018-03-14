package commands

import (
	"encoding/json"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/models"
	"github.com/spf13/cobra"
)

var createStatsCmd = &cobra.Command{
	Use:   "stats [resource]",
	Short: "Create stats for a resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateStats(args[0], cmd.Flag("statsjson").Value.String(), client)
	},
}

func init() {
	createStatsCmd.Flags().String("statsjson", "", "JSON-encoded stats to be uploaded")
	createStatsCmd.MarkFlagRequired("statsjson")
	createCmd.AddCommand(createStatsCmd)
}

func CreateStats(resource string, statsJson string, client clients.VRopsClientIntf) error {
	stats := []models.Stat{}
	err := json.Unmarshal([]byte(statsJson), &stats)
	if err != nil {
		return err
	}
	err = client.CreateStats(resource, stats)
	if err != nil {
		return err
	}
	return nil
}
