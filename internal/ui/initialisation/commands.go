package initialisation

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/fs"
	tea "github.com/charmbracelet/bubbletea"
	"time"
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
			return progressMsg(fmt.Sprintf("%s", step.displayName))
		})

		cmds = append(cmds, func() tea.Msg {
			err := step.run()
			if err != nil {
				return progressMsg(fmt.Sprintf(" %s failed: %v", step.displayName, err))
			}
			return progressMsg(fmt.Sprintf("%s completed", step.displayName))
		})
	}

	cmds = append(cmds, func() tea.Msg {
		return progressMsg("Setup complete!")
	})

	return tea.Sequence(cmds...)
}

func createDirCmd(name, path string) tea.Cmd {
	return func() tea.Msg {
		_, err := fs.EnsureDirExists(path)
		status := "created"
		if err != nil {
			status = "failed"
		}
		time.Sleep(5 * time.Second)
		return initResultMsg{
			Name:   name,
			Result: status,
			Err:    err,
		}
	}
}
