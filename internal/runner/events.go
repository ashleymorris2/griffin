package runner

type Event interface{}

type TaskStarted struct {
	StepLabel string
	TaskLabel string
}

type TaskOutput struct {
	TaskLabel string
	Content   string
}

type TaskFinished struct {
	TaskLabel string
}

type TaskFailed struct {
	TaskLabel string
	Err       error
}
