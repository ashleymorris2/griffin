package menu

import (
	"github.com/ashleymorris2/booty/internal/ui/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	list     list.Model
	quitting bool
	done     bool
	Result   string
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
