package listselect

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
)

type SelectorItem struct {
	TitleText       string
	DescriptionText string
	Value           string
}

func (s SelectorItem) Title() string       { return s.TitleText }
func (s SelectorItem) Description() string { return s.DescriptionText }
func (s SelectorItem) FilterValue() string { return s.TitleText }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SelectorItem)
	if !ok {
		return
	}

	isSelected := index == m.Index()
	cursor := " "
	title := i.TitleText
	desc := i.DescriptionText

	if isSelected {
		cursor = styles.Cursor
		title = styles.HighlightedListItem.Render(title)
		desc = styles.HighlightedListDescription.Render(desc)
	} else {
		title = styles.ListItem.Render(title)
		desc = styles.ListDescription.Render(desc)
	}

	fmt.Fprintf(w, "%s %s\n", cursor, title)
	fmt.Fprintf(w, "%s %s\n", cursor, desc)
}

type ListSelectorModel struct {
	list     list.Model
	quitting bool
	done     bool
	Result   string
}

func New(title string, items []SelectorItem) ListSelectorModel {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, itemDelegate{}, 40, 10)
	l.SetFilteringEnabled(true)
	l.SetShowPagination(true)
	l.Title = title
	l.Styles.Title = styles.Title
	l.Styles.PaginationStyle = styles.Pagination
	l.Styles.HelpStyle = styles.HelpStyle

	return ListSelectorModel{list: l}
}
