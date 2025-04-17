package seqtask

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type taskStatusType int

const (
	statusPending taskStatusType = iota
	statusInProgress
	statusSuccess
	statusFailed
)

type taskStatus struct {
	Status  taskStatusType
	Message string
}

// progressMsg is sent to the update loop when a task has progressed in executeTaskAsync
type progressMsg struct {
	stepId string
	status taskStatus
}

// executeTaskAsync launches a background goroutine that runs a single SequentialTask.
// It sends progressMsg status updates (pending → in-progress → success/failure) to the provided channel.
//
// This design keeps the UI responsive by offloading long-running operations
// to a separate goroutine and communicating results back via progressMsg values.
func executeTaskAsync(task SequentialTask, ch chan progressMsg) {
	go func() {
		// Send "pending" status immediately
		ch <- progressMsg{
			task.ID,
			taskStatus{Status: statusPending, Message: task.Message + " (pending)"},
		}
		time.Sleep(time.Duration(rand.Int63n(250)+100) * time.Millisecond)

		// Send "in progress" status after a short delay
		ch <- progressMsg{
			task.ID,
			taskStatus{Status: statusInProgress, Message: task.Message},
		}
		time.Sleep(time.Duration(rand.Int63n(500)+100) * time.Millisecond)

		// Execute the task
		result, err := task.Run()
		if err != nil {
			time.Sleep(time.Duration(rand.Int63n(200)+100) * time.Millisecond)

			// Report failure
			ch <- progressMsg{
				task.ID,
				taskStatus{Status: statusFailed, Message: fmt.Sprintf("%s - %v", task.Message, err)},
			}

			return
		}
		time.Sleep(time.Duration(rand.Int63n(200)+100) * time.Millisecond)

		// Report success
		ch <- progressMsg{
			task.ID,
			taskStatus{Status: statusSuccess, Message: result},
		}
	}()
}

// waitForTaskProgress returns a command that waits for the next progress message.
// It runs in the background and doesn’t block the UI.
// When a message is received, Bubble Tea passes it to Update.
// If no message arrives, it just keeps waiting.
func waitForTaskProgress(ch chan progressMsg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}
