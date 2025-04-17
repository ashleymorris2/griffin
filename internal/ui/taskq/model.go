package taskq

import "github.com/charmbracelet/bubbles/spinner"

type SequentialTask struct {
	ID      string
	Message string
	Run     func() (string, error)
}

type SequentialTaskModel struct {
	tasks       []SequentialTask      // The tasks that are to be Run
	taskChan    chan progressMsg      // Channel used to send progress updates between task runners and the model
	statuses    map[string]taskStatus // Tracks the status (pending, success, failure) for each task by ID
	taskCount   int                   // Total number of tasks to Run
	currentTask int                   // Index of the current task being executed
	spinner     spinner.Model         // Spinner UI element from the Bubble Tea bubbles package
	finished    bool                  // Flag indicating whether all tasks have completed
}

func NewSequentialTaskModel(tasks []SequentialTask) SequentialTaskModel {

	s := spinner.New()
	s.Spinner = spinner.Jump

	statuses := make(map[string]taskStatus)
	for _, step := range tasks {
		statuses[step.ID] = taskStatus{
			Status:  statusPending,
			Message: step.Message + " (queued)",
		}
	}

	return SequentialTaskModel{
		tasks:       tasks,
		spinner:     s,
		taskCount:   len(tasks),
		statuses:    statuses,
		taskChan:    make(chan progressMsg),
		currentTask: 0,
	}
}

// nextTask returns the next initialization task to Run, the updated model, and a boolean
// indicating whether a task was available.
//
// The model is passed and returned by value, which is intentional. In Bubble Tea,
// models are typically treated as immutable: rather than modifying state in place,
// changes are returned as a new version of the model.
//
// This approach helps avoid unintended side effects and keeps state transitions explicit.
func (m SequentialTaskModel) nextTask() (SequentialTaskModel, SequentialTask, bool) {
	if m.currentTask >= len(m.tasks) {
		return m, SequentialTask{}, false
	}

	task := m.tasks[m.currentTask]
	m.currentTask++
	return m, task, true
}
