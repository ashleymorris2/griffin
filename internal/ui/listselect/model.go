package listselect

import "github.com/charmbracelet/bubbles/list"

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

func NewListSelectorModel(title string, items []SelectorItem) ListSelectorModel {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, list.NewDefaultDelegate(), 40, 10)
	l.Title = title
	return ListSelectorModel{list: l}
}
