package runner

import "github.com/ashleymorris2/booty/internal/modules"

type StepRunner struct {
	modules     map[string]modules.Module // registered modules (e.g., shell, docker)
	stopOnError bool                      // configuration: stop blueprint if a task fails
}

func NewStepRunner(modules map[string]modules.Module, stopOnError bool) *StepRunner {
	return &StepRunner{
		modules:     modules,
		stopOnError: stopOnError,
	}
}
