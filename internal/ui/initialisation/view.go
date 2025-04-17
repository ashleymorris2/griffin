package initialisation

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type progressMsg struct {
	stepId string
	status taskStatus
}

func (m initModel) Init() tea.Cmd {

	// Start the first step
	executeTaskAsync(m.tasks[m.currentTask], m.taskChan)

	// Wait for the first progress message
	return tea.Batch(
		m.spinner.Tick,
		waitForStepProgress(m.taskChan),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case progressMsg:
		m.statuses[msg.stepId] = msg.status

		// Check if the step has completed
		if msg.status.Status == statusSuccess || msg.status.Status == statusFailed {
			m, task, ok := m.nextTask()
			if ok {
				executeTaskAsync(task, m.taskChan)
				return m, waitForStepProgress(m.taskChan)
			} else {
				// No more steps to run
				m.finished = true
				return m, nil
			}
		}

		// If still in-progress, just continue ticking
		return m, waitForStepProgress(m.taskChan)

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
		b.WriteString("Initialization complete!\n")
	} else {
		b.WriteString(fmt.Sprintf("Running initalization...\n"))
	}

	for _, stepID := range taskOrder {
		if step, ok := m.statuses[stepID]; ok {
			var symbol string
			switch step.Status {
			case statusSuccess:
				symbol = "[✓]"
			case statusFailed:
				symbol = "[X]"
			case statusInProgress:
				symbol = fmt.Sprintf("[%s]", m.spinner.View())
			default:
				symbol = "[ ]"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", symbol, step.Message))
		}
	}

	result := b.String()
	return result
}
