package initialisation

import "github.com/ashleymorris2/booty/internal/fs"

func PrepareSetupFolder(path string) (string, error) {
	status, err := fs.EnsurePathExistsInHome(path)
	switch status {
	case fs.StatusCreated:
		return "Environment ready.", nil
	case fs.StatusAlreadyExists:
		return "Environment already exists. Skipping step.", nil
	case fs.StatusFailed:
		return "Failed to prepare environment.", err
	default:
		return "Unknown result during environment setup.", err
	}
}
