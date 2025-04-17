package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

type PathStatus int

const (
	StatusCreated PathStatus = iota
	StatusAlreadyExists
	StatusFailed
)

type PathCheckResult struct {
	Status PathStatus
	Path   string
}

// EnsureSubdirInHome checks for the existence of a subdirectory within the user's home directory.
// If the directory does not exist, it attempts to create it with permission 0700.
//
// The function returns a PathCheckResult containing the status (e.g. created, already exists, or failed)
// and the full absolute path to the directory. If any error occurs (e.g. home directory not found,
// permission denied), the returned error will contain additional context.
//
// This is useful for safely creating app-specific folders in a cross-platform way under $HOME (Linux/macOS)
// or %USERPROFILE% (Windows).
func EnsureSubdirInHome(subdir string) (PathCheckResult, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return PathCheckResult{StatusFailed, home}, fmt.Errorf("unable to find home directory: %w", err)
	}

	fullPath := filepath.Join(home, subdir)

	// Check if the directory exists
	info, err := os.Stat(fullPath)

	switch {
	case os.IsNotExist(err):
		// Try to create (read/write/execute - owner only)
		if err := os.Mkdir(fullPath, 0700); err != nil {
			return PathCheckResult{
				Status: StatusFailed,
				Path:   "",
			}, fmt.Errorf("failed to create directory: %w", err)
		}
		return PathCheckResult{
			Status: StatusCreated,
			Path:   fullPath,
		}, nil
	case err == nil && info.IsDir():
		// Directory exists
		return PathCheckResult{
			Status: StatusAlreadyExists,
			Path:   fullPath,
		}, nil
	case err != nil:
		return PathCheckResult{
			Status: StatusFailed,
			Path:   "",
		}, fmt.Errorf("error checking directory: %w", err)
	default:
		// Edge case: path exists but is not a directory
		return PathCheckResult{
			Status: StatusFailed,
			Path:   "",
		}, fmt.Errorf("path exists but is not a directory: %s", fullPath)
	}
}

// WriteFileToHomeSubdir writes the given file contents to a file with the specified name
// inside a subdirectory of the user's home directory. If the subdirectory does not exist,
// it will be created automatically.
//
// For example, calling WriteFileToHomeSubdir("myapp", "config.yaml", data) will create or
// ensure the directory $HOME/myapp exists and write the contents of data to $HOME/myapp/config.yaml.
//
// Returns an error if the subdirectory cannot be created or if the file write fails.
func WriteFileToHomeSubdir(subDir string, filename string, file []byte) (string, error) {
	result, err := EnsureSubdirInHome(subDir)
	if err != nil {
		return "", fmt.Errorf("failed to find or create destination directory: %w", err)
	}

	destPath := filepath.Join(result.Path, filename)

	// Check if file already exists
	if _, err := os.Stat(destPath); err == nil {
		return "File already exists. (skipped)", nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("error checking file: %w", err)
	}

	// (File permissions (read/write - owner, read - group)
	err = os.WriteFile(destPath, file, 0640)
	err = os.WriteFile(destPath, file, 0640)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return "File created successfully.", nil
}
