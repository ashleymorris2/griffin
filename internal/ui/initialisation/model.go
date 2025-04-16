package initialisation

import "github.com/charmbracelet/bubbles/spinner"

type initModel struct {
	statuses       map[string]stepProgress
	totalSteps     int
	completedSteps int
	spinner        spinner.Model
	finished       bool
}

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Jump

	return initModel{
		spinner:    s,
		totalSteps: len(stepOrder),
		statuses:   make(map[string]stepProgress),
	}
}
