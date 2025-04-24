package pick

import (
	"errors"
	"fmt"
	"github.com/ashleymorris2/booty/internal/core/blueprint"
	"github.com/ashleymorris2/booty/internal/ui/components/menu"
	tea "github.com/charmbracelet/bubbletea"
)

func BlueprintFrom(files []string) (string, error) {

	var results = blueprint.ReadMetadataFromFiles(files)

	items := make([]menu.Item, len(files))
	for res := range results {
		if res.Err != nil {
			continue
		}
		items[res.Index] = menu.NewItem(res.Item.Title, res.Item.Description, res.Item.FilePath)
	}

	program := tea.NewProgram(menu.New("Choose a configuration to run", items))
	res, err := program.Run()
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	if m, ok := res.(menu.Model); ok && m.Result != "" {
		return m.Result, nil
	}

	return "", errors.New("no result")
}
