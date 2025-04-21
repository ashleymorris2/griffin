package cmd

import (
	"github.com/ashleymorris2/booty/internal/core/configselect"
	"github.com/ashleymorris2/booty/internal/fs"
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
	Run: func(cmd *cobra.Command, args []string) {
		RunInteractiveSelector()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func RunInteractiveSelector() {
	files, err := fs.ListFilesInSubDirectory("config")
	err = configselect.Run(files)
	if err != nil {
		return
	}
}
