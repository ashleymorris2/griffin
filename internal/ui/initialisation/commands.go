package initialisation

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type initStep struct {
	id      string
	message string
	run     func() (string, error)
}

func buildSetupCommands(steps []initStep) tea.Cmd {
	var cmds []tea.Cmd

	for _, step := range steps {
		step := step
		cmds = append(cmds, makeStepCmd(step))
	}

	return tea.Sequence(cmds...)
}

func makeStepCmd(step initStep) tea.Cmd {
	return tea.Sequence(
		func() tea.Msg {
			return progressMsg{
				stepId: step.id,
				status: stepProgress{
					Status:  statusPending,
					Message: step.message,
				},
			}
		},
		func() tea.Msg {
			time.Sleep(300 * time.Millisecond)
			return progressMsg{
				stepId: step.id,
				status: stepProgress{
					Status:  statusInProgress,
					Message: step.message,
				},
			}
		},
		func() tea.Msg {
			time.Sleep(1 * time.Second) // Simulate a delay - just because
			result, err := step.run()
			if err != nil {
				return progressMsg{
					stepId: step.id,
					status: stepProgress{
						Status:  statusFailed,
						Message: fmt.Sprintf("%s - %v", result, err),
					},
				}
			}
			return progressMsg{
				stepId: step.id,
				status: stepProgress{
					Status:  statusSuccess,
					Message: result,
				},
			}
		},
	)
}
