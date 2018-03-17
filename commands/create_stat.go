package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/models"
	"github.com/spf13/cobra"
)

var (
	statKey, statTime string
	value             float64
)
var createStatCmd = &cobra.Command{
	Use:   "stat [adapterkind]",
	Short: "Create a single statistic for a resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdErr = CreateStat(args[0], statKey, statTime, value, client)
	},
}

func init() {
	createStatCmd.Flags().StringVarP(&statKey, "stat-key", "s", "", "stat key for the stat (in format key|subkey)")
	createStatCmd.Flags().StringVar(&statTime, "time", "", "time to use (if blank, current time is used). Can be a timestamp in ms or time in RFC3339 format")
	createStatCmd.Flags().Float64VarP(&value, "value", "v", 0, "Value of the statistic")
	createStatCmd.MarkFlagRequired("stat-key")
	createStatCmd.MarkFlagRequired("value")
	createCmd.AddCommand(createStatCmd)
}

func CreateStat(resource, statKey, statTime string, value float64, client clients.VRopsClientIntf) error {
	var timestamp int64
	if statTime == "" {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	} else if ts, ok := strconv.ParseInt(statTime, 10, 64); ok == nil {
		timestamp = ts
	} else {
		t1, e := time.Parse(time.RFC3339, statTime)
		if e != nil {
			return fmt.Errorf("Cannot parse time: %s", e)
		}
		timestamp = t1.UnixNano() / int64(time.Millisecond)
	}

	stats := models.Stat{
		Data:       []float64{value},
		Timestamps: []int64{timestamp},
		StatKey:    statKey,
	}

	err := client.CreateStats(resource, []models.Stat{stats})
	if err != nil {
		return err
	}
	return nil
}
