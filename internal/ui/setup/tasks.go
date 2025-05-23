package setup

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/files"
	"github.com/ashleymorris2/booty/internal/fs"
	"path/filepath"
)

const setupFolderPath = ".devsetup"
const (
	stepPrepareEnv    = "prepare-env"           // Set up the environment (e.g., ensure required dirs exist)
	stepCreateExample = "create-example-config" // Create an example config file (disable with --example=false
)

func tasks() []SequentialTask {
	return []SequentialTask{
		{
			ID:      stepPrepareEnv,
			Message: "Preparing local environment...",
			Run: func() (string, error) {
				resultMsg, err := prepareLocalEnvironment()
				return resultMsg, err
			},
		},
		{
			ID:      stepCreateExample,
			Message: "Creating example config file...",
			Run: func() (string, error) {
				result, err := createExampleConfig()
				return result, err
			},
		},
	}
}

func prepareLocalEnvironment() (string, error) {
	result, err := fs.EnsureSubdirectoryExists(setupFolderPath)
	if err != nil {
		return "", fmt.Errorf("failed to prepare environment: %w", err)
	}

	switch result.Status {
	case fs.StatusCreated:
		return "Environment ready.", nil
	case fs.StatusAlreadyExists:
		return "Environment already prepared. (skipped)", nil
	default:
		return "", fmt.Errorf("unknown result during environment setup %w", err)
	}
}

func createExampleConfig() (string, error) {
	result, err := fs.WriteToSubdirectory(filepath.Join(setupFolderPath, "config"), "example.yml", files.ExampleConfig)
	if err != nil {
		return "", fmt.Errorf("could not write config file: %w", err)
	}

	return result, nil
}
