package utils

import (
	"envsafes/utils"
	"os"
)

func DeleteTokenFile() error {
	// Implement this function to delete the token file
	err := os.Remove(utils.GetExecutablePath() + "/auth.token")
	if err != nil {
		return err
	}
	return nil
}
