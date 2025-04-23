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

func (d itemDelegate) Height() int                             { return 2 }
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
		desc = styles.HighlightedListDesc.Render(desc)
	} else {
		title = styles.ListItem.Render(title)
		desc = styles.ListDesc.Render(desc)
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

	l := list.New(listItems, itemDelegate{}, 40, 20)

	l.Title = title

	l.Styles.Title = styles.Title

	l.FilterInput.PromptStyle = styles.FilterPrompt

	l.Help.Styles.ShortKey = styles.Help
	l.Help.Styles.ShortDesc = styles.HelpDesc
	l.Help.Styles.FullKey = styles.Help
	l.Help.Styles.FullDesc = styles.HelpDesc

	l.Styles.PaginationStyle = styles.Pagination

	return ListSelectorModel{list: l}
}
