package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	failMsg    = "Environment setup failed."
	createdMsg = "Environment ready."
	skipMsg    = "Environment already exists. Skipping step."
)

type PathStatus int

const (
	StatusCreated PathStatus = iota
	StatusAlreadyExists
	StatusFailed
)

// EnsurePathExistsInHome Returns true if a directory exists with the given name in $HOME on Linux and macOS or
// %USERPROFILE% on Windows, and attempts to create it if it doesn't.
func EnsurePathExistsInHome(path string) (PathStatus, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return StatusFailed, fmt.Errorf("unable to find home directory: %w", err)
	}

	fullPath := filepath.Join(home, path)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := os.Mkdir(fullPath, 0700); err != nil {
			return StatusFailed, fmt.Errorf("failed to create directory: %w", err)
		}
		return StatusCreated, nil
	} else if err == nil {
		return StatusAlreadyExists, nil
	} else {
		return StatusFailed, fmt.Errorf("error checking directory: %w", err)
	}
}
