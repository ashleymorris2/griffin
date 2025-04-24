package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var PrimaryColor = lipgloss.Color("#87FFAF")
var SecondaryColor = lipgloss.Color("#5eb27a")

var itemColor = lipgloss.Color("15")
var descColor = lipgloss.Color("250")
var highlightItemColor = PrimaryColor
var highlightDescColor = SecondaryColor

type ListStyles struct {
	Cursor          string
	Title           lipgloss.Style
	Item            lipgloss.Style
	Description     lipgloss.Style
	Highlighted     lipgloss.Style
	HighlightedDesc lipgloss.Style
	Pagination      lipgloss.Style
	Help            lipgloss.Style
	HelpDesc        lipgloss.Style
	FilterPrompt    lipgloss.Style
}

var List = ListStyles{
	Cursor:          lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true).SetString("┃").String(),
	Title:           lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true),
	Item:            lipgloss.NewStyle().Foreground(itemColor),
	Description:     lipgloss.NewStyle().Foreground(descColor).Faint(true),
	Highlighted:     lipgloss.NewStyle().Foreground(highlightItemColor).Bold(true),
	HighlightedDesc: lipgloss.NewStyle().Foreground(highlightDescColor),
	Pagination:      list.DefaultStyles().PaginationStyle.PaddingLeft(4).UnsetForeground(),
	Help:            list.DefaultStyles().HelpStyle.Foreground(SecondaryColor),
	HelpDesc:        list.DefaultStyles().HelpStyle.Foreground(lipgloss.Color("247")),
	FilterPrompt:    lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true),
}
