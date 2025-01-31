package cmd

import (
	"envsafes/utils"
	localutils "envsafes/utils/local_env"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [command]",
	Short:   "Run the command with the environment variables",
	Example: `envsafes run python script.py`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// No arguments provided, show help
			cmd.Help()
			os.Exit(0)
		}
		var restArgs []string
		restArgs = append(restArgs, args[1:]...)

		run := exec.Command(args[0], restArgs...)
		run.Env = os.Environ()

		if config := utils.GetConfig(); config.EnvLocation == "local" {
			// Load local environment variables
			localEnvVars, err := localutils.GetLocalEnvVars()
			if err != nil {
				log.Fatalf("Error loading local environment variables: %v", err)
			}

			run.Env = append(run.Env, localEnvVars...)

		} else {
			// load env from remote server
		}

		// Set up output to the console
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr

		if err := run.Run(); err != nil {
			log.Fatalf("Execution failed: %v", err)
		}
	},
}
