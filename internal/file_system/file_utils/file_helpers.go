package fileutils

import (
	"os"
	"fmt"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		// File or directory exists
		return true
	}

	if os.IsNotExist(err) {
		// File or directory does not exist
		return false
	}

	// An error occurred while accessing the file or directory
	fmt.Printf("Error checking file: %v\n", err)
	return false
}