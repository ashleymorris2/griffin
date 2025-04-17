package initialisation

import "github.com/charmbracelet/bubbles/spinner"

type initModel struct {
	tasks       []initTask            // The tasks that are to be run
	taskChan    chan progressMsg      // Channel used to send progress updates between task runners and the model
	statuses    map[string]taskStatus // Tracks the status (pending, success, failure) for each task by ID
	taskCount   int                   // Total number of tasks to run
	currentTask int                   // Index of the current task being executed
	spinner     spinner.Model         // Spinner UI element from the Bubble Tea bubbles package
	finished    bool                  // Flag indicating whether all tasks have completed
}

func newInitModel() initModel {
	tasks := buildSteps()

	s := spinner.New()
	s.Spinner = spinner.Jump

	statuses := make(map[string]taskStatus)
	for _, step := range tasks {
		statuses[step.id] = taskStatus{
			Status:  statusPending,
			Message: step.message + " (queued)",
		}
	}

	return initModel{
		tasks:       tasks,
		spinner:     s,
		taskCount:   len(taskOrder),
		statuses:    statuses,
		taskChan:    make(chan progressMsg),
		currentTask: 0,
	}
}

// nextTask returns the next initialization task to run, the updated model, and a boolean
// indicating whether a task was available.
//
// The model is passed and returned by value, which is intentional. In Bubble Tea,
// models are typically treated as immutable: rather than modifying state in place,
// changes are returned as a new version of the model.
//
// This approach helps avoid unintended side effects and keeps state transitions explicit.
func (m initModel) nextTask() (initModel, initTask, bool) {
	if m.currentTask >= len(m.tasks) {
		return m, initTask{}, false
	}

	task := m.tasks[m.currentTask]
	m.currentTask++
	return m, task, true
}
