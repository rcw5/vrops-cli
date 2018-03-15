package commands

import (
	"encoding/json"

	"github.com/rcw5/vrops-cli/clients"
	"github.com/rcw5/vrops-cli/models"
	"github.com/spf13/cobra"
)

var createResourceCmd = &cobra.Command{
	Use: "resource",
	Long: `Create a new resource and associated adapterkind and resourcekind, if they do not already exist.

See samples/resource.json for an example payload.`,
	Short: "Create a resource (and associated adapterkind)",
	Run: func(cmd *cobra.Command, args []string) {
		cmdErr = CreateResource(cmd.Flag("definition").Value.String(), client)
	},
}

func init() {
	createResourceCmd.Flags().String("definition", "", "JSON-encoded definition of the new resource (required)")
	createResourceCmd.MarkFlagRequired("definition")
	createCmd.AddCommand(createResourceCmd)
}

func CreateResource(definition string, client clients.VRopsClientIntf) error {
	resource := models.Resource{}
	err := json.Unmarshal([]byte(definition), &resource)
	if err != nil {
		return err
	}
	err = client.CreateResource(resource)
	if err != nil {
		return err
	}
	return nil
}
