package messages

// Sent when a task starts running
type TaskStartedMsg struct {
	StepLabel string
	TaskLabel string
}

// Sent when a task emits a line of output
type TaskOutputMsg struct {
	TaskLabel string
	Content   string
}

// Sent when a task finishes successfully
type TaskFinishedMsg struct {
	TaskLabel string
}

// Sent when a task fails with an error
type TaskFailedMsg struct {
	TaskLabel string
	Err       error
}
