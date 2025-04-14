package ui

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/core"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

lURL   string
}

type initModel struct {
	dependencies []core.Dependency
	current      int
	statuses     []string
	spinner      spinner.Model
	finished     bool
}

type tickMsg struct{}

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return initModel{
		dependencies: []core.Dependency{
			{Name: "WSL", CheckCommand: "wsl"},
		},
		spinner: s,
	}
}

func (m initModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return tickMsg{}
		}),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.finished || m.current >= len(m.dependencies) {
			m.finished = true
			return m, tea.Quit
		}

		dependency := &m.dependencies[m.current]

		switch dependency.Status {
		case core.Initial:
			dependency.Status = core.Checking
			return m, tickAfter(500 * time.Millisecond)
		case core.Checking:
			dependency.Check()


			// Schedule next tick after short delay
			return m, tickAfter(500 * time.Millisecond)
		case core.Installing:
			//todo run install logic here
			dependency.Status = done
			m.current++
			return m, tickAfter(500 * time.Millisecond)
		case core.Done:
			return m, tea.Quit
		}
	}

	// Done with all steps?
	if m.current >= len(m.dependencies) {
		m.finished = true
		return m, tea.Quit
	}

	return m, nil
}

func (m initModel) View() string {
	if m.finished {
		out := "\nSummary:\n"
		for _, d := range m.dependencies {
			out += renderOutput(d) + "\n"
		}
		out += "\nPress Ctrl+C to exit.\n"
		return out
	}

	out := "\n"

	for _, d := range m.dependencies {
		out += renderOutput(d) + "\n"
	}

	return out
}

func RunBootstrapUI() error {
	fmt.Println("Initialising your environment...")

	p := tea.NewProgram(newInitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		return err
	}

	return nil
}

func tickAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func renderOutput(d core.Dependency) string {
	switch d.Status {
	case core.Done:
		if d.Found {
			return fmt.Sprintf("✅ %s found", d.Name)
		}
		return fmt.Sprintf("⚠️  %s not found. Installed manually or skipped.", d.Name)
	case checking:
		return fmt.Sprintf("🔍 Checking for %s...", d.Name)
	case installing:
		return fmt.Sprintf("⬇️  Install required for %s → %s", d.Name, d.InstallURL)
	default:
		return fmt.Sprintf("⏺️  Waiting to check %s", d.Name)
	}
}
