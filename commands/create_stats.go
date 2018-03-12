package commands

import (
	"encoding/json"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/models"
)

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
