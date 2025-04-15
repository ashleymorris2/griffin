package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

// EnsureDirExists Returns true if a directory exists with the given name in $HOME on Linux and macOS or
// %USERPROFILE% on Windows, and attempts to create it if it doesn't.
func EnsureDirExists(path string) (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	setupDir := filepath.Join(home, path)
	if _, err := os.Stat(setupDir); os.IsNotExist(err) {
		err := os.Mkdir(setupDir, 0700)
		if err != nil {
			return false, fmt.Errorf("failed to create setup directory: %w", err)
		}
	}

	return true, nil
}
