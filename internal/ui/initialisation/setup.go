package initialisation

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/files"
	"github.com/ashleymorris2/booty/internal/fs"
)

type stepStatusType int

const (
	statusPending stepStatusType = iota
	statusInProgress
	statusSuccess
	statusFailed
)

const setupFolderPath = ".devsetup"

type stepProgress struct {
	Status  stepStatusType
	Message string
}

func prepareLocalEnvironment() (string, error) {
	result, err := fs.EnsureSubdirInHome(setupFolderPath)
	if err != nil {
		return "", fmt.Errorf("failed to prepare environment: %w", err)
	}

	switch result.Status {
	case fs.StatusCreated:
		return "Environment ready.", nil
	case fs.StatusAlreadyExists:
		return "Environment already exists. Skipping step.", nil
	default:
		return "", fmt.Errorf("unknown result during environment setup %w", err)
	}
}

func createExampleConfig() error {
	err := fs.WriteFileToHomeSubdir("config", "example.yml", files.ExampleConfig)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}
