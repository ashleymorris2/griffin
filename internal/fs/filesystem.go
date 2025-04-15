package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirStatus int

const (
	Created DirStatus = iota
	Errored
	Exists
)

// EnsureDirExists Returns true if a directory exists with the given name in $HOME on Linux and macOS or
// %USERPROFILE% on Windows, and attempts to create it if it doesn't.
func EnsureDirExists(path string) (DirStatus, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Errored, err
	}

	setupDir := filepath.Join(home, path)
	if _, err := os.Stat(setupDir); os.IsNotExist(err) {
		err := os.Mkdir(setupDir, 0700)
		if err != nil {
			return Errored, fmt.Errorf("failed to create setup directory: %w", err)
		}
	} else if err == nil {
		return Exists, nil
	} else {
		return Errored, fmt.Errorf("error checking directory: %w", err)
	}

	return Created, nil
}
