package initialisation

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/files"
	"github.com/ashleymorris2/booty/internal/fs"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"path/filepath"
	"time"
)

type taskStatusType int

const (
	statusPending taskStatusType = iota
	statusInProgress
	statusSuccess
	statusFailed
)

const setupFolderPath = ".devsetup"

type taskStatus struct {
	Status  taskStatusType
	Message string
}

// executeTaskAsync launches a background goroutine that runs a single init task.
// It sends progressMsg updates (queued → in-progress → success/failure) to the provided channel.
//
// This design keeps the UI responsive by offloading long-running operations
// to a separate goroutine and communicating results back via progressMsg values.
func executeTaskAsync(task initTask, ch chan progressMsg) {
	go func() {
		// Send "pending" status immediately
		ch <- progressMsg{
			task.id,
			taskStatus{Status: statusPending, Message: task.message + " (pending)"},
		}
		time.Sleep(time.Duration(rand.Int63n(250)+100) * time.Millisecond)

		// Send "in progress" status after a short delay
		ch <- progressMsg{
			task.id,
			taskStatus{Status: statusInProgress, Message: task.message},
		}
		time.Sleep(time.Duration(rand.Int63n(500)+100) * time.Millisecond)

		// Execute the task
		result, err := task.run()
		if err != nil {
			time.Sleep(time.Duration(rand.Int63n(200)+100) * time.Millisecond)

			// Report failure
			ch <- progressMsg{
				task.id,
				taskStatus{Status: statusFailed, Message: fmt.Sprintf("%s - %v", task.message, err)},
			}

			return
		}
		time.Sleep(time.Duration(rand.Int63n(200)+100) * time.Millisecond)

		// Report success
		ch <- progressMsg{
			task.id,
			taskStatus{Status: statusSuccess, Message: result},
		}
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
	err := fs.WriteFileToHomeSubdir(filepath.Join(setupFolderPath, "config"), "example.yml", files.ExampleConfig)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}
