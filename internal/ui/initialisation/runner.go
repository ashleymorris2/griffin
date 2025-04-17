package initialisation

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/files"
	"github.com/ashleymorris2/booty/internal/fs"
	tea "github.com/charmbracelet/bubbletea"
	"time"
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

func runStep(step initStep, ch chan progressMsg) {
	go func() {
		ch <- progressMsg{step.id, stepProgress{Status: statusPending, Message: step.message + " (queued)"}}
		time.Sleep(400 * time.Millisecond)

		ch <- progressMsg{step.id, stepProgress{Status: statusInProgress, Message: step.message}}
		time.Sleep(600 * time.Millisecond)

		result, err := step.run()
		if err != nil {
			time.Sleep(300 * time.Millisecond)
			ch <- progressMsg{step.id, stepProgress{Status: statusFailed, Message: fmt.Sprintf("%s - %v", step.message, err)}}
			return
		}
		time.Sleep(300 * time.Millisecond)
		ch <- progressMsg{step.id, stepProgress{Status: statusSuccess, Message: result}}
	}()
}

func waitForStepProgress(ch chan progressMsg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
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
