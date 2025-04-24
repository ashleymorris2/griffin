package pick

import (
	"errors"
	"fmt"
	"github.com/ashleymorris2/booty/internal/core/blueprint"

	"github.com/ashleymorris2/booty/internal/ui/components/listselect"
	tea "github.com/charmbracelet/bubbletea"
)

func BlueprintFrom(files []string) (string, error) {

	metadata, _ := blueprint.ReadMetadataFromFiles(files)

	items := make([]listselect.SelectorItem, len(metadata))
	for i, data := range metadata {
		items[i] = listselect.SelectorItem{
			TitleText:       data.Title,
			DescriptionText: data.Description,
			Value:           data.Path,
		}
	}

	model := listselect.New("Choose a configuration to run", items)
	program := tea.NewProgram(model)
	result, err := program.Run()
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	if m, ok := result.(listselect.ListSelectorModel); ok && m.Result != "" {
		return m.Result, nil
	}

	return "", errors.New("no result")
}
