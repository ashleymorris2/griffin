package cmd

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/fs"
	"github.com/ashleymorris2/booty/internal/modules"
	"github.com/ashleymorris2/booty/internal/runner"
	"github.com/ashleymorris2/booty/internal/ui/components/taskrunner"
	"github.com/ashleymorris2/booty/internal/ui/messages"
	"github.com/ashleymorris2/booty/internal/ui/pick"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := selectBlueprintPath()
		if err != nil {
			return err
		}

		err = runBlueprint(path)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func selectBlueprintPath() (string, error) {
	files, err := fs.ListFilesInSubDirectory("config")
	if err != nil {
		return "", fmt.Errorf("error listing files: %v\\n\"", err)
	}
	if len(files) == 0 {
		return "", fmt.Errorf("no files found in 'config' directory")
	}

	path, err := pick.BlueprintFrom(files)
	if err != nil {
		return "", fmt.Errorf("error selecting runner %w", err)
	}

	fmt.Println("Selected:", path)

	return path, nil
}

func runBlueprint(path string) error {
	// 1. Load blueprint
	bp, err := fs.ReadBlueprintFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to read blueprint: %w", err)
	}

	// 2. Extract task labels (flatten all steps)
	var taskLabels []string
	for _, step := range bp.Steps {
		for _, task := range step.Tasks {
			taskLabels = append(taskLabels, task.Label)
		}
	}

	// 3. Create TUI model
	model := taskrunner.NewAccordionModel(taskLabels)
	program := tea.NewProgram(model)

	// 4. Create event -> tea.Msg mapper
	emit := func(ev runner.Event) {
		switch e := ev.(type) {
		case runner.TaskStarted:
			program.Send(messages.TaskStartedMsg{StepLabel: e.StepLabel, TaskLabel: e.TaskLabel})
		case runner.TaskOutput:
			program.Send(messages.TaskOutputMsg{TaskLabel: e.TaskLabel, Content: e.Content})
		case runner.TaskFinished:
			program.Send(messages.TaskFinishedMsg{TaskLabel: e.TaskLabel})
		case runner.TaskFailed:
			program.Send(messages.TaskFailedMsg{TaskLabel: e.TaskLabel, Err: e.Err})
		}
	}

	// 5. Register modules with OnOutput callback
	mod := modules.Register(func(line string) {
		emit(runner.TaskOutput{TaskLabel: model.Tasks[model.CurrentIndex], Content: line})
	})

	// 6. Create and start runner
	r := runner.New(mod, emit, false, 10)
	go func() {
		_ = r.RunBlueprint(bp)
	}()

	// 7. Run the TUI
	_, err = program.Run()
	return err
}
