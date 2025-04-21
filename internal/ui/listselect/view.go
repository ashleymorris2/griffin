package listselect

import tea "github.com/charmbracelet/bubbletea"

type ItemSelectedMsg struct {
	Value string
}

func (m ListSelectorModel) Init() tea.Cmd {
	return nil
}

func (m ListSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.list.SelectedItem().(SelectorItem); ok {
				return m, func() tea.Msg {
					return ItemSelectedMsg{Value: item.Value}
				}
			}
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListSelectorModel) View() string {
	if m.quitting {
		return "Cancelled.\n"
	}
	return m.list.View()
}
