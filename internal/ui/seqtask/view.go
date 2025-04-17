package seqtask

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type initMsg struct{}

func (m SequentialTaskModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg { return initMsg{} },
	)
}

func (m SequentialTaskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case initMsg:
		mm, nextTask, ok := m.getNextTask()
		if ok {
			executeTaskAsync(nextTask, mm.taskChan)
			return mm, waitForTaskProgress(mm.taskChan)
		}
		return m, nil
	case progressMsg:
		m.statuses[msg.stepId] = msg.status

		// Check if the step has completed
		if msg.status.Status == statusSuccess || msg.status.Status == statusFailed {
			mm, nextTask, ok := m.getNextTask()
			if ok {
				// Execute the next getNextTask and wait for progress
				executeTaskAsync(nextTask, mm.taskChan)
				return mm, waitForTaskProgress(mm.taskChan)
			} else {
				// No more steps to Run
				mm.finished = true
				return mm, nil
			}
		}

		// If still in-progress, just continue ticking
		return m, waitForTaskProgress(m.taskChan)

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

func (m SequentialTaskModel) View() string {
	var b strings.Builder

	if m.finished {
		b.WriteString("Initialization complete\n\n")
	} else {
		b.WriteString(fmt.Sprintf("Running initalization...\n\n"))
	}

	for _, task := range m.tasks {
		if step, ok := m.statuses[task.ID]; ok {
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
