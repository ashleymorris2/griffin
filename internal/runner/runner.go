package runner

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/models"
	"github.com/ashleymorris2/booty/internal/modules"
	"sync"
)

type Runner struct {
	modules     *modules.ModuleRegistry // All the available modules
	stopOnError bool                    // configuration: stop runner if a task fails
	maxWorkers  int
}

func New(modules *modules.ModuleRegistry, stopOnError bool, maxWorkers int) *Runner {
	//if modules == nil {
	//	modules = make(map[string]modules.)
	//}
	return &Runner{
		modules:     modules,
		stopOnError: stopOnError,
		maxWorkers:  maxWorkers,
	}
}

func (r *Runner) RunBlueprint(bp *models.Blueprint) error {
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

func (r *Runner) runTasks(tasks []models.Task) {
	// Cap workers between 1 - len(tasks)
	workers := len(tasks)
	if r.maxWorkers > 0 && r.maxWorkers < workers {
		workers = r.maxWorkers
	}

	taskChan := make(chan models.Task)
	errChan := make(chan error, 1)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskChan {
				if module, ok := r.modules.Get(task.Uses); ok {
					if err := module.Run(task); err != nil {
						errChan <- fmt.Errorf("task %s by worker %d failed: %w", task.Uses, workerID, err)
					}
				} else {
					errChan <- fmt.Errorf("worker %d: no module found for task: %s", workerID, task.Uses)
				}
			}
		}(i)
	}

	// Send tasks to workers
	go func() {
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()

	wg.Wait() // Wait for the workers to finish
}
