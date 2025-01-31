package utils

import (
	"fmt"
)

func RemoveLocalEnvVar() error {
	// Remove a local environment variable
	envVars, err := GetLocalEnvVars()
	if err != nil {
		return fmt.Errorf("error loading local environment variables: %v", err)
	}
	if len(envVars) == 0 {
		fmt.Println("No environment variables found.")
		return nil
	}

	for i, envVar := range envVars {
		fmt.Printf("(%d) %s\n", i, envVar)
	}
label:
	fmt.Print("Enter the number of the environment variable to remove: ")
	var index int
	fmt.Scanln(&index)

	if index < 0 || index >= len(envVars) {
		fmt.Println("invalid index")
		goto label
	}

	removedEnvVar := envVars[index]
	envVars = append(envVars[:index], envVars[index+1:]...)
	line := ""
	for i, envVar := range envVars {
		if i == len(envVars)-1 {
			line += envVar
			break
		}
		line += envVar + "\n"
	}
	if err := WriteToLocalEnv(line); err != nil {
		return fmt.Errorf("error writing to local environment file: %v", err)
	}
	fmt.Printf("Environment variable '%s' removed successfully.\n", removedEnvVar)
	return nil
}
