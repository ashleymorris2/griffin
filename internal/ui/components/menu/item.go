package menu

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
)

type Item struct {
	title       string
	description string
	Value       string
}

func NewItem(title, description, value string) Item {
	return Item{
		title:       title,
		description: description,
		Value:       value,
	}
}

func (s Item) FilterValue() string { return s.title }

type itemDelegate struct{}

func (d itemDelegate) Height() int { return 2 }

func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	isSelected := index == m.Index()

	cursor := " "
	title := i.title
	desc := i.description

	if isSelected {
		cursor = styles.List.Cursor
		title = styles.List.Highlighted.Render(title)
		desc = styles.List.HighlightedDesc.Render(desc)
	} else {
		title = styles.List.Item.Render(title)
		desc = styles.List.Description.Render(desc)
	}

	_, _ = fmt.Fprintf(w, "%s %s\n", cursor, title)
	_, _ = fmt.Fprintf(w, "%s %s\n", cursor, desc)
}
