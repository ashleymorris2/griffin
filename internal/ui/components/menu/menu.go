package menu

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func Show(title string, items []Item) (Model, error) {
	p := tea.NewProgram(New(title, items))
	res, err := p.Run()
	if err != nil {
		return Model{}, fmt.Errorf("error: %s", err)
	}

	if m, ok := res.(Model); ok {
		return m, nil
	}

	return Model{}, errors.New("unexpected result: model type mismatch")
}
