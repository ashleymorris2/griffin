package runner

import (
	"github.com/ashleymorris2/booty/internal/models"
	"github.com/ashleymorris2/booty/internal/modules"
	"sync"
)

type Runner struct {
	modules     map[string]modules.Module // All the available modules
	stopOnError bool                      // configuration: stop blueprint if a task fails
	maxWorkers  int
}

func (r *Runner) RunBlueprint(bp models.Blueprint) error {
	for _, step := range bp.Steps {
		err := r.runStep(step)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) runStep(s models.Step) error {
	r.runTasks(s.Tasks)

	return nil
}

func (r *Runner) runTasks(t []models.Task) {
	jobs := make([]models.Task, 0, len(t))
	errChan := make(chan error, 1)
	var wg sync.WaitGroup

	for i := 0; i < r.maxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
			}
		}()
	}
}

//func (r *Runner) runTask(t models.Task) error {
//
//}
