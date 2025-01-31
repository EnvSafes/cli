package cmd

import (
	utils "envsafes/utils/auth"
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from EnvSafe",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.DeleteTokenFile()
		if err != nil {
			fmt.Println("Failed to logout")
			return
		}
		fmt.Println("Logged out successfully")
	},
}
