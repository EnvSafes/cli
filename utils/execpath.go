package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		os.Exit(1)
	}

	// Convert it to an absolute path (if needed)
	absExePath, err := filepath.Abs(exePath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}
	return filepath.Dir(absExePath)
}
