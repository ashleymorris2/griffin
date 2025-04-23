package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var primaryColor = lipgloss.Color("#87FFAF")
var accentColor = lipgloss.Color("#5eb27a")

var (
	Cursor              = lipgloss.NewStyle().Foreground(primaryColor).Bold(true).SetString("┃").String()
	Title               = lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	ListItem            = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	ListDesc            = lipgloss.NewStyle().Foreground(lipgloss.Color("247")).Faint(true)
	HighlightedListItem = lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	HighlightedListDesc = lipgloss.NewStyle().Foreground(accentColor)
	Pagination          = list.DefaultStyles().PaginationStyle.PaddingLeft(4).UnsetForeground()
	Help                = list.DefaultStyles().HelpStyle.Foreground(accentColor)
	HelpDesc            = list.DefaultStyles().HelpStyle.Foreground(lipgloss.Color("247"))
	FilterPrompt        = lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	QuitText            = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
