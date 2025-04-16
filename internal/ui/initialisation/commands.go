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
			return progressMsg{step.id, step.message}
		})

		cmds = append(cmds, func() tea.Msg {
			result, err := step.run()
			if err != nil {
				return progressMsg{stepId: step.id, message: fmt.Sprintf(" %s - %v", result, err)}
			}
			return progressMsg{step.id, result}
		})
	}

	//cmds = append(cmds, func() tea.Msg {
	//	return progressMsg("Setup complete!")
	//})

	return tea.Sequence(cmds...)
}
