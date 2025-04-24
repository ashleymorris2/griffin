package taskrunner

import "github.com/charmbracelet/bubbles/spinner"

type SequentialTask struct {
	ID      string
	Message string
	Run     func() (string, error)
}

type SequentialTaskRunnerModel struct {
	tasks        []SequentialTask      // The tasks that are to be run
	taskChan     chan progressMsg      // Channel used to send progress updates between task runners and the model
	statuses     map[string]taskStatus // Tracks the status (pending, success, failure) for each task by ID
	taskCount    int                   // Total number of tasks to Run
	currentTask  int                   // Index of the current task being executed
	spinner      spinner.Model         // Spinner UI element from the Bubble Tea bubbles package
	finished     bool                  // Flag indicating whether all tasks have completed
	initialTitle string                // Title to display when the program begins
	finalTitle   string                // Title to display when the program ends
}

func New(tasks []SequentialTask, initialTitle string, finalTitle string) SequentialTaskRunnerModel {

	s := spinner.New()
	s.Spinner = spinner.Jump

	statuses := make(map[string]taskStatus)
	for _, step := range tasks {
		statuses[step.ID] = taskStatus{
			Status:  statusPending,
			Message: step.Message + "(queued)",
		}
	}

	return SequentialTaskRunnerModel{
		tasks:        tasks,
		taskChan:     make(chan progressMsg),
		statuses:     statuses,
		taskCount:    len(tasks),
		currentTask:  0,
		spinner:      s,
		finished:     false,
		initialTitle: initialTitle,
		finalTitle:   finalTitle,
	}
}

// getNextTask returns the next setup task to Run, the updated model, and a boolean
// indicating whether a task was available.
//
// The model is passed and returned by value, which is intentional. In Bubble Tea,
// models are typically treated as immutable: rather than modifying state in place,
// changes are returned as a new version of the model.
//
// This approach helps avoid unintended side effects and keeps state transitions explicit.
func (m SequentialTaskRunnerModel) getNextTask() (SequentialTaskRunnerModel, SequentialTask, bool) {
	if m.currentTask >= len(m.tasks) {
		return m, SequentialTask{}, false
	}

	task := m.tasks[m.currentTask]
	m.currentTask++
	return m, task, true
}
