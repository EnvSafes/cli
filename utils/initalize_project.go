package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	EnvLocation string `json:"env_location"`
	ProjectID   string `json:"project_id,omitempty"`
	ProjectName string `json:"project_name,omitempty"`
}

func InitializeProject() error {
	var envLocation string

	fmt.Print("Enter environment location (local/remote): ")
	fmt.Scanln(&envLocation)

	if envLocation != "local" && envLocation != "remote" {
		return fmt.Errorf("invalid environment location, must be 'local' or 'remote'")
	}
	// Initialize config
	config := Config{
		EnvLocation: envLocation,
	}
	if envLocation == "remote" {
		fmt.Print("Enter project ID: ")
		fmt.Scanln(&config.ProjectID)
	} else {
		fmt.Print("Enter project name: ")
		fmt.Scanln(&config.ProjectName)
		envPath := GetExecutablePath() + "/env_vars"
		err := os.MkdirAll(envPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", envPath, err)
		}
		filePath := envPath + "/" + config.ProjectName + ".env"
		_, err = os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %v", filePath, err)
		}
	}

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Write JSON data to file
	fileName := "envsafes.config.json"
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Config file '%s' created successfully.\n", fileName)
	return nil
}
