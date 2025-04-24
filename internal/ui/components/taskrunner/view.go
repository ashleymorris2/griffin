package taskrunner

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type initCompleteMsg struct{}

func (m SequentialTaskRunnerModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg { return initCompleteMsg{} },
	)
}

func (m SequentialTaskRunnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case initCompleteMsg:
		mm, nextTask, ok := m.getNextTask()
		if ok {
			executeTask(nextTask, mm.taskChan)
			return mm, waitForTaskProgress(mm.taskChan)
		}
		// No tasks - set as finished
		mm.finished = true
		return mm, nil

	case progressMsg:
		m.statuses[msg.taskId] = msg.status

		// Check if the step has completed
		if msg.status.Status == statusSuccess || msg.status.Status == statusFailed {
			mm, nextTask, ok := m.getNextTask()
			if ok {
				// Execute the next getNextTask and wait for progress
				executeTask(nextTask, mm.taskChan)
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

func (m SequentialTaskRunnerModel) View() string {

	var b strings.Builder

	title := m.initialTitle
	if m.finished {
		title = m.finalTitle
	}

	b.WriteString(fmt.Sprintf("%s\n\n", title))

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
