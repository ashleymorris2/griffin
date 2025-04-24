package menu

import (
	"errors"
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui/styles"
	"github.com/charmbracelet/bubbles/key"
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

func (d itemDelegate) Height() int                             { return 2 }
func (d itemDelegate) Spacing() int                            { return 0 }
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

type Model struct {
	list     list.Model
	quitting bool
	done     bool
	Result   string
}

func Show(items []Item) (Model, error) {
	p := tea.NewProgram(New("Choose a configuration to run", items))
	res, err := p.Run()
	if err != nil {
		return Model{}, fmt.Errorf("error: %s", err)
	}

	if m, ok := res.(Model); ok {
		return m, nil
	}

	return Model{}, errors.New("unexpected result: model type mismatch")
}

func New(title string, items []Item) Model {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, itemDelegate{}, 40, 20)

	l.Title = title
	l.Styles.Title = styles.List.Title

	l.FilterInput.PromptStyle = styles.List.FilterPrompt

	l.Help.Styles.ShortKey = styles.List.Help
	l.Help.Styles.ShortDesc = styles.List.HelpDesc
	l.Help.Styles.FullKey = styles.List.Help
	l.Help.Styles.FullDesc = styles.List.HelpDesc

	l.Styles.PaginationStyle = styles.List.Pagination

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Enter,
		}
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Enter,
		}
	}

	return Model{list: l}
}
