package commands

import (
	"fmt"
	"os"

	"github.com/rcw5/vrops-cli/clients"

	"github.com/spf13/cobra"
)

var client clients.VRopsClientIntf

var rootCmd = &cobra.Command{
	Use:   "vrops-cli",
	Short: "A straightforward way to interact with the vROps API",
	Long:  `Wouldn't it be amazing if you didn't need to worry about the vROps REST API!?`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		username := envVarFlagOrFail("VROPS_USERNAME", "username", cmd)
		password := envVarFlagOrFail("VROPS_PASSWORD", "password", cmd)
		url := envVarFlagOrFail("VROPS_URL", "target", cmd)
		client = clients.NewVROpsClient(url, username, password, trace)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve stuff from vROps",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Change stuff in vROps",
}

var trace bool

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "vROps username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "vROps password")
	rootCmd.PersistentFlags().StringP("target", "t", "", "url to vROps instance")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "enable request tracing")

	getCmd.PersistentFlags().StringP("output", "o", "table", "Output format (table or json)")

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func envVarFlagOrFail(envVarName, flagName string, cmd *cobra.Command) string {
	value := os.Getenv(envVarName)
	if value == "" {
		value = cmd.Flag(flagName).Value.String()
		if value == "" {
			fmt.Printf("ERROR: You must set either $%s or --%s\n", envVarName, flagName)
			os.Exit(1)
		}
	}
	return value
}
