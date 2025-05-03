package runner

import (
	"fmt"
	"sync"

	"github.com/ashleymorris2/booty/internal/models"
	"github.com/ashleymorris2/booty/internal/modules"
)

type Runner struct {
	modules     *modules.ModuleRegistry // All the available modules
	stopOnError bool                    // Configuration: stop runner if a task fails
	maxWorkers  int
	emit        func(Event)
}

func New(modules *modules.ModuleRegistry, emit func(Event), stopOnError bool, maxWorkers int) *Runner {
	// if modules == nil {
	//	modules = make(map[string]modules.)
	// }
	return &Runner{
		modules:     modules,
		stopOnError: stopOnError,
		maxWorkers:  maxWorkers,
		emit:        emit,
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

	r.startWorkers(workers, taskChan, errChan, &wg)
	r.dispatchTasks(tasks, taskChan)

	wg.Wait() // Wait for the workers to finish
}

// startWorkers starts a pool of count 'workers', accepts a readonly channel to read tasks from,
// a write only channel to send errors to and a WaitGroup to signal that workers have completed all available tasks
func (r *Runner) startWorkers(workers int, taskChan <-chan models.Task, errChan chan<- error, wg *sync.WaitGroup) {
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskChan {
				if module, ok := r.modules.Get(task.Uses); ok {
					r.emit(TaskStarted{StepLabel: task.Label, TaskLabel: task.Label})
					if err := module.Run(task); err != nil {
						errChan <- fmt.Errorf("task %s by worker %d failed: %w", task.Uses, workerID, err)
					} else {
						r.emit(TaskFinished{TaskLabel: task.Label})
					}
				} else {
					errChan <- fmt.Errorf("worker %d: no module found for task: %s", workerID, task.Uses)
				}
			}
		}(i)
	}
}

func (r *Runner) dispatchTasks(tasks []models.Task, taskChan chan<- models.Task) {
	go func() {
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()
}
