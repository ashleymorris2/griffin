package configselect

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/ui/listselect"
	"gopkg.in/yaml.v3"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type configFileMetadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func Run(files []string) error {

	var items []listselect.SelectorItem
	for _, file := range files {
		meta, err := readMetadata(file)
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
		return fmt.Errorf("error: %s", err)
	}

	if msg, ok := result.(listselect.ListSelectorModel); ok {
		fmt.Println("Running:", msg)
		// runSetupFromFile(msg.Value)
	}

	return nil
}

func readMetadata(path string) (configFileMetadata, error) {
	var meta configFileMetadata

	data, err := os.ReadFile(path)
	if err != nil {
		return meta, err
	}

	err = yaml.Unmarshal(data, &meta)
	return meta, err
}
