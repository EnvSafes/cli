package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfig() Config {
	var config Config

	// Read config file
	file, err := os.ReadFile("envsafes.config.json")
	if err != nil {
		fmt.Println("Error reading config file")
		return Config{}
	}

	// Unmarshal JSON data
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data")
		return Config{}
	}

	return config
}
