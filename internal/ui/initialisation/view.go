package initialisation

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type progressMsg struct {
	stepId  string
	message string
}

func (m initModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		buildSetupCommands([]initStep{
			{
				id:      stepPrepareEnv,
				message: "Preparing local environment...",
				run: func() (string, error) {
					resultMsg, err := PrepareSetupFolder(".devsetup")
					return resultMsg, err
				},
			},
		}),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case progressMsg:
		m.statuses[msg.stepId] = msg.message
		return m, nil

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

	//if m.finished {
	//	b.WriteString("Setup complete!")
	//	b.WriteString(strings.Join(m.statuses, "\n"))
	//	b.WriteString("Nice")
	//} else {
	//	b.WriteString(fmt.Sprintf("%s Setting up...\n%s", m.spinner.View(), strings.Join(m.statuses, "\n")))
	//}

	result := b.String()
	return result
}
