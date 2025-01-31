package utils

import (
	"fmt"
	"strings"
)

func AddLocalEnvVar() error {

	// Get the key and value from the user
	var key string
	var value string
label:
	fmt.Print("Enter the key: ")
	fmt.Scanln(&key)
	fmt.Print("Enter the value: ")
	fmt.Scanln(&value)

	if key == "" || value == "" {
		fmt.Println("key and value cannot be empty")
		goto label
	}
	envVars, err := GetLocalEnvVars()
	if err != nil {
		return fmt.Errorf("error loading local environment variables: %v", err)
	}
	envVars = append(envVars, fmt.Sprintf("%s=%s", strings.TrimSpace(key), strings.TrimSpace(value)))
	line := ""
	for i, envVar := range envVars {
		if i == len(envVars)-1 {
			line += envVar
			break
		}
		line += envVar + "\n"
	}

	return WriteToLocalEnv(line)
}
