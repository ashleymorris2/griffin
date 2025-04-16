package initialisation

import "github.com/ashleymorris2/booty/internal/fs"

type stepStatusType int

const (
	statusPending stepStatusType = iota
	statusInProgress
	statusSuccess
	statusFailed
)

type stepProgress struct {
	Status  stepStatusType
	Message string
}

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
