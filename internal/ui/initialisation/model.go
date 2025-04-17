package initialisation

import "github.com/charmbracelet/bubbles/spinner"

type initModel struct {
	stepChan    chan progressMsg
	statuses    map[string]stepProgress
	totalSteps  int
	currentStep int
	spinner     spinner.Model
	finished    bool
}

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Jump

	return initModel{
		spinner:     s,
		totalSteps:  len(stepOrder),
		statuses:    make(map[string]stepProgress),
		stepChan:    make(chan progressMsg),
		currentStep: 0,
	}
}
