package initialisation

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type initStep struct {
	displayName string
	run         func() error
}

func buildSetupCommands(steps []initStep) tea.Cmd {
	var cmds []tea.Cmd

	for _, step := range steps {
		step := step
		cmds = append(cmds, func() tea.Msg {
			return progressMsg(fmt.Sprintf("[%s] - Running", step.displayName))
		})

		cmds = append(cmds, func() tea.Msg {
			err := step.run()
			if err != nil {
				return progressMsg(fmt.Sprintf(" [%s] - failed: %v", step.displayName, err))
			}
			return progressMsg(fmt.Sprintf("[%s] - Completed", step.displayName))
		})
	}

	cmds = append(cmds, func() tea.Msg {
		return progressMsg("Setup complete!")
	})

	return tea.Sequence(cmds...)
}
