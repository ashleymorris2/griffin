package cmd

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui/listselect"
	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"

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

type SetupMetadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func RunInteractiveSelector() {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".devsetup", "config")

	files, _ := filepath.Glob(filepath.Join(configDir, "*.yml"))
	if len(files) == 0 {
		fmt.Println("No setup files found.")
		return
	}

	// Convert filenames to SelectorItems
	var items []listselect.SelectorItem
	for _, file := range files {
		meta, err := ReadMetadata(file)
		if err != nil {
			continue
		}

		items = append(items, listselect.SelectorItem{
			TitleText:       meta.Title,
			DescriptionText: meta.Description,
			Value:           file,
		})
	}

	model := listselect.New("Choose a configuration to run", items)
	program := tea.NewProgram(model)
	result, err := program.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if msg, ok := result.(listselect.ListSelectorModel); ok {
		fmt.Println("Running:", msg)
		// Now run your actual setup logic here
		// runSetupFromFile(msg.Value)
	}
}

func ReadMetadata(path string) (SetupMetadata, error) {
	var meta SetupMetadata

	data, err := os.ReadFile(path)
	if err != nil {
		return meta, err
	}

	err = yaml.Unmarshal(data, &meta)
	return meta, err
}
