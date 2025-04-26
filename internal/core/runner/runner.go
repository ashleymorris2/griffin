package runner

import (
	"github.com/ashleymorris2/booty/internal/core/blueprint"
	"github.com/ashleymorris2/booty/internal/modules"
)

type Runner struct {
	modules     map[string]modules.Module // registered modules (e.g., shell, docker)
	stopOnError bool                      // configuration: stop blueprint if a task fails
}

func (r *Runner) RunBlueprint(bp blueprint.Blueprint) error {
	for _, step := range bp.Steps {
		runner := NewStepRunner(r.modules, r.stopOnError)
		err := StepRunner.RunStep(step)
		if err != nil {
			return err
		}
	}
	return nil
}
