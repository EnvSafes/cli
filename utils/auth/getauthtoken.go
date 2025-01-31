package utils

import (
	"envsafes/utils"
	"os"
)

func GetAuthToken() string {
	// Implement this function to return the access token
	execPath := utils.GetExecutablePath()
	file, err := os.Open(execPath + "/auth.token")
	if err != nil {
		return ""
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		return ""
	}
	return string(buf[:n])
}
