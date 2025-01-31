package cmd

import (
	utils "envsafes/utils/auth"
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to EnvSafe",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement login logic here
		utils.Authenticate()
		fmt.Println("Logged in successfully")
	},
}
