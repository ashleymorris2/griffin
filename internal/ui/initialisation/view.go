package initialisation

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/fs"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type progressMsg string

type initResultMsg struct {
	Name   string
	Result string
	Err    error
}

func (m initModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		buildSetupCommands([]initStep{
			{
				displayName: "Create .devsetup folder",
				run: func() error {
					_, err := fs.EnsureDirExists(".devsetup")
					return err
				},
			},
		}),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case progressMsg:
		newModel := m
		newModel.statuses = append(newModel.statuses, string(msg))
		return newModel, nil

	case initResultMsg:
		newModel := m
		newModel.statuses = append(newModel.statuses, fmt.Sprintf("%s: %s", msg.Name, msg.Result))
		//m.statuses = append(m.statuses, fmt.Sprintf("%s: %s", msg.Name, msg.Result))
		if msg.Err != nil {
			fmt.Println("Error:", msg.Err)
		}
		newModel.finished = true
		return newModel, tea.Quit

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m initModel) View() string {
	var b strings.Builder

	if m.finished {
		b.WriteString("Setup complete!")
		b.WriteString(fmt.Sprintf("%s", strings.Join(m.statuses, "\n")))
		b.WriteString("Nice")
	} else {
		b.WriteString(fmt.Sprintf("%s Setting up...\n%s", m.spinner.View(), strings.Join(m.statuses, "\n")))
	}

	return b.String()
}
