package initialisation

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type progressMsg struct {
	stepId string
	status stepProgress
}

func (m initModel) Init() tea.Cmd {

	for _, step := range buildSteps() {
		m.statuses[step.id] = stepProgress{
			Status:  statusPending,
			Message: step.message + " (queued)",
		}
	}

	// Start the first step
	runStep(buildSteps()[m.currentStep], m.stepChan)

	// Wait for the first progress message
	return tea.Batch(
		m.spinner.Tick,
		waitForStepProgress(m.stepChan),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case progressMsg:
		m.statuses[msg.stepId] = msg.status

		// Check if the step has completed
		if msg.status.Status == statusSuccess || msg.status.Status == statusFailed {
			m.currentStep++
			if m.currentStep < m.totalSteps {
				runStep(buildSteps()[m.currentStep], m.stepChan)
			} else {
				return m, tea.Quit
			}
		}
		return m, waitForStepProgress(m.stepChan)

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

	for _, stepID := range stepOrder {
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
