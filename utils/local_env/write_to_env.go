package utils

import (
	"envsafes/utils"
	"fmt"
	"os"
)

func WriteToLocalEnv(line string) error {
	config := utils.GetConfig()
	envPath := utils.GetExecutablePath() + "/env_vars"
	filePath := envPath + "/" + config.ProjectName + ".env"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening config file")
	}
	if _, err := file.WriteString(line); err != nil {
		return fmt.Errorf("error writing to config file")
	}
	return nil
}
