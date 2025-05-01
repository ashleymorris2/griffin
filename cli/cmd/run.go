package cmd

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/fs"
	"github.com/ashleymorris2/booty/internal/modules"
	"github.com/ashleymorris2/booty/internal/runner"
	"github.com/ashleymorris2/booty/internal/ui/pick"
	"github.com/spf13/cobra"
	"os"
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
	Run: func(cmd *cobra.Command, args []string) {
		path, err := selectBlueprintPath()
		if err != nil {
			return
		}

		err = runBlueprint(path)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func selectBlueprintPath() (string, error) {
	files, err := fs.ListFilesInSubDirectory("config")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error listing files: %v\n", err)
		return "", err
	}
	if len(files) == 0 {
		return "", fmt.Errorf("no files found in 'config' directory")
	}

	path, err := pick.BlueprintFrom(files)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error selecting runner: %v\n", err)
		return "", err
	}

	fmt.Println("Selected:", path)

	return path, nil
}

func runBlueprint(path string) error {
	bp, err := fs.ReadBlueprintFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to read blueprint: %w", err)
	}

	m := modules.RegisterModules()
	r := runner.New(m, false, 10)

	err = r.RunBlueprint(bp)
	if err != nil {
		return err
	}
	return nil
}
