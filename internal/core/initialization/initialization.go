package initialization

import (
	"fmt"

	"github.com/ashleymorris2/booty/internal/ui/seqtask"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(
		seqtask.New(
			tasks(),
			"Initialization running... hang tight 😎",
			"Initialization complete 😌",
		),
	)

	_, err := p.Run()
	if err != nil {
		fmt.Printf("There's been an error: %v", err)
		return err
	}

	return nil
}
