package initialisation

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(newInitModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("There's been an error: %v", err)
		return err
	}

	// Printing finalModel.View() keeps the final view on screen when exiting the program,
	// otherwise it gets cleared
	fmt.Println(finalModel.View())

	return nil
}
