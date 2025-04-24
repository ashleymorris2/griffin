package menu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ItemSelectedMsg struct {
	Value string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Enter):
			if item, ok := m.list.SelectedItem().(Item); ok {
				m.done = true
				m.Result = item.Value
				return m, tea.Quit
			}
		case msg.String() == "q" || msg.String() == "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return "Cancelled.\n"
	}
	return m.list.View()
}
