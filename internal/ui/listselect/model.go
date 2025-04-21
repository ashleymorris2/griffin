package listselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

type SelectorItem struct {
	TitleText       string
	DescriptionText string
	Value           string // Can be used to store an ID, filename, etc.
}

func (s SelectorItem) Title() string       { return s.TitleText }
func (s SelectorItem) Description() string { return s.DescriptionText }
func (s SelectorItem) FilterValue() string { return s.TitleText }

type ListSelectorModel struct {
	list     list.Model
	quitting bool
	done     bool
	result   string
}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(0)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("15"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(40).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

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
	cursor := "  "
	title := i.TitleText
	desc := i.DescriptionText

	if isSelected {
		cursor = "> "
		title = selectedItemStyle.Render(title)
		desc = selectedItemStyle.Render(desc)
	} else {
		title = itemStyle.Render(title)
		desc = itemStyle.Render(desc)
	}

	fmt.Fprintf(w, "%s%s\n", cursor, title)
	fmt.Fprintf(w, "  %s\n", desc)
}
func New(title string, items []SelectorItem) ListSelectorModel {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, itemDelegate{}, 40, 10)
	l.SetFilteringEnabled(true)
	l.Title = title
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return ListSelectorModel{list: l}
}
