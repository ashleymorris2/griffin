package menu

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Enter key.Binding
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "select"),
	),
}
