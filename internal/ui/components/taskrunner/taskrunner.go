package taskrunner

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// internal/ui/runner/accordion_model.go

type AccordionModel struct {
	Expanded     map[string]bool     // taskLabel -> expanded
	Output       map[string][]string // taskLabel -> output lines
	Tasks        []string            // ordered task labels
	CurrentIndex int                 // current selection index
	Status       map[string]string   // taskLabel -> "pending", "done", "failed"
}

func NewAccordionModel(taskLabels []string) AccordionModel {
	return AccordionModel{
		Expanded:     make(map[string]bool),
		Output:       make(map[string][]string),
		Tasks:        taskLabels,
		CurrentIndex: 0,
		Status:       make(map[string]string),
	}
}

func (m AccordionModel) Init() tea.Cmd {
	return nil
}

func (m AccordionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.CurrentIndex > 0 {
				m.CurrentIndex--
			}
		case "down":
			if m.CurrentIndex < len(m.Tasks)-1 {
				m.CurrentIndex++
			}
		case "enter", " ":
			label := m.Tasks[m.CurrentIndex]
			m.Expanded[label] = !m.Expanded[label]
		}

	case ui.TaskStartedMsg:
		m.Status[msg.TaskLabel] = "running"
	case ui.TaskOutputMsg:
		m.Output[msg.TaskLabel] = append(m.Output[msg.TaskLabel], msg.Content)
	case ui.TaskFinishedMsg:
		m.Status[msg.TaskLabel] = "done"
	case ui.TaskFailedMsg:
		m.Status[msg.TaskLabel] = "failed"
	}

	return m, nil
}

func (m AccordionModel) View() string {
	var b strings.Builder
	for i, label := range m.Tasks {
		prefix := "  "
		if i == m.CurrentIndex {
			prefix = "> "
		}
		status := m.Status[label]
		if status == "done" {
			status = "✓"
		} else if status == "failed" {
			status = "✗"
		} else if status == "running" {
			status = "…"
		}
		b.WriteString(fmt.Sprintf("%s[%s] %s\n", prefix, status, label))

		if m.Expanded[label] {
			for _, line := range m.Output[label] {
				b.WriteString("    " + line + "\n")
			}
		}
	}
	return b.String()
}
