package cmd

import (
	utils "envsafes/utils/local_env"

	"github.com/spf13/cobra"
)

var localEnvCMD = &cobra.Command{
	Use:   "local",
	Short: "Work with local environment variables",
}

func init() {
	localEnvCMD.AddCommand(addLocalEnvCMD)
	localEnvCMD.AddCommand(listLocalEnvCMD)
	localEnvCMD.AddCommand(removeLocalEnvCMD)
}

var addLocalEnvCMD = &cobra.Command{
	Use:   "add",
	Short: "Add a new local environment variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.AddLocalEnvVar()
	},
}

var listLocalEnvCMD = &cobra.Command{
	Use:   "list",
	Short: "List all local environment variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.ListLocalEnvVars()
	},
}

var removeLocalEnvCMD = &cobra.Command{
	Use:   "remove",
	Short: "Remove a local environment variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.RemoveLocalEnvVar()
	},
}
