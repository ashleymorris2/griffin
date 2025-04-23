package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var primaryColor = lipgloss.Color("#87FFAF")
var accentColor = lipgloss.Color("#5eb27a")

// List specific styles
var (
	Cursor                     = lipgloss.NewStyle().Foreground(primaryColor).Bold(true).SetString("┃").String()
	Title                      = lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	ListItem                   = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	ListDescription            = lipgloss.NewStyle().Foreground(lipgloss.Color("247")).Faint(true)
	HighlightedListItem        = lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	HighlightedListDescription = lipgloss.NewStyle().Foreground(accentColor)
	Pagination                 = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle                  = list.DefaultStyles().HelpStyle.PaddingLeft(3).PaddingBottom(1)
	QuitText                   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
