package initialisation

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
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
		cmds = append(cmds, func() tea.Msg {
			return progressMsg{stepId: step.id, status: stepProgress{Status: statusInProgress, Message: step.message}}
		})

		cmds = append(cmds, func() tea.Msg {
			result, err := step.run()
			if err != nil {
				message := fmt.Sprintf(" %s - %v", result, err)
				return progressMsg{stepId: step.id, status: stepProgress{Status: statusFailed, Message: message}}
			}
			return progressMsg{stepId: step.id, status: stepProgress{Status: statusSuccess, Message: result}}
		})
	}

	cmds = append(cmds, func() tea.Msg {
		return initCompleteMsg("Setup complete!")
	})

	return tea.Sequence(cmds...)
}
