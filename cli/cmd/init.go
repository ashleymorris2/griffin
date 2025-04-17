package cmd

import (
	"github.com/ashleymorris2/booty/internal/tasks/initialization"
	"github.com/spf13/cobra"
)

// initCmd represents the taskq command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialises ",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := initialization.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
