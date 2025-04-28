package runner

import (
	"github.com/ashleymorris2/booty/internal/models"
	"github.com/ashleymorris2/booty/internal/modules"
)

type Runner struct {
	modules     map[string]modules.Module // registered modules (e.g., shell, docker)
	stopOnError bool                      // configuration: stop blueprint if a task fails
}

func (r *Runner) RunBlueprint(bp models.Blueprint) error {
	//step := newStepRunner(r.modules, r.stopOnError)

	for _, step := range bp.Steps {

		err := step.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
