package initialisation

import "github.com/charmbracelet/bubbles/spinner"

type initModel struct {
	statuses map[string]stepProgress
	spinner  spinner.Model
	finished bool
}

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return initModel{
		spinner:  s,
		statuses: make(map[string]stepProgress),
	}
}
