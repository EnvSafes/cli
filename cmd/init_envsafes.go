package cmd

import (
	"envsafes/utils"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the envsafes environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.InitializeProject()
	},
}
