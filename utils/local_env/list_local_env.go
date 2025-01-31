package utils

import "fmt"

func ListLocalEnvVars() error {
	// List all local environment variables
	envVar, err := GetLocalEnvVars()
	if err != nil {
		return fmt.Errorf("error getting local environment variables: %v", err)
	}

	if len(envVar) == 0 {
		fmt.Println("No local environment variables found")
		return nil
	}
	for i, v := range envVar {
		fmt.Println("(", i, ") ", v)
	}

	return nil
}
