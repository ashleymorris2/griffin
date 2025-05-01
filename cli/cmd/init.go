package cmd

import (
	"github.com/ashleymorris2/booty/internal/ui/setup"
	"github.com/spf13/cobra"
)

type InitOptions struct {
	SkipExample bool
}

var initOpts InitOptions

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prepares the required folder structure and config files",
	Long: `Initializes your local development environment by creating a ".devsetup" folder
in your $HOME directory with user-only permissions. 

Configuration files are generated and managed exclusively in this location to avoid tampering or accidental exposure.

Re-run this command anytime to recreate or verify your setup.

Example:
  booty init`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := setup.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&initOpts.SkipExample, "no-example", false, "Skip generating the example runner")
}
