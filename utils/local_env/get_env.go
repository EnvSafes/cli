package utils

import (
	"envsafes/utils"
	"fmt"
	"os"
	"strings"
)

func GetLocalEnvVars() ([]string, error) {
	// Get local environment variables
	config := utils.GetConfig()
	if config.EnvLocation != "local" {
		return nil, fmt.Errorf("invalid config: %s, it must be local", config.EnvLocation)
	}
	var envVars []string

	// Read local environment variables
	path := utils.GetExecutablePath()

	// Read the local environment variables
	file, err := os.ReadFile(path + "/env_vars/" + config.ProjectName + ".env")
	if err != nil {
		return nil, fmt.Errorf("env variables doesnt exist please add them using \"envsafes local add\" command")
	}

	envVars = strings.Split(string(file), "\n")
	if len(envVars) == 1 && envVars[0] == "" {
		envVars = []string{}
	}
	//skip empty lines
	for i := 0; i < len(envVars); i++ {
		if strings.TrimSpace(envVars[i]) == "" {
			envVars = append(envVars[:i], envVars[i+1:]...)
			i--
		}
	}
	return envVars, nil
}
