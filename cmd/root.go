package cmd

import (
	"envsafes/utils"
	authutils "envsafes/utils/auth"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envsafes",
	Short: "EnvSafes is a CLI tool for managing environment variables and running commands with them securely",
	Long:  `EnvSafes is a CLI tool that allows you to securely manage and organize environment variables in groups.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if token := authutils.GetAuthToken(); token != "" {
		rootCmd.AddCommand(logoutCmd)
	} else {
		rootCmd.AddCommand(loginCmd)
	}
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
	if config := utils.GetConfig(); config.EnvLocation == "local" {
		rootCmd.AddCommand(localEnvCMD)
	}

}
