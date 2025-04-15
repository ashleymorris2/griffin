package ui

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/core"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type dependenciesModel struct {
	dependencies []core.Dependency
	current      int
	statuses     []string
	spinner      spinner.Model
	finished     bool
}

type tickMsg struct{}

func newDependenciesModel() dependenciesModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return dependenciesModel{
		dependencies: []core.Dependency{
			{Name: "WSL", CheckCommand: "wsl"},
		},
		spinner: s,
	}
}

func (m dependenciesModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tea.Tick(time.Millisecond*400, func(t time.Time) tea.Msg {
			return tickMsg{}
		}),
	)
}

func (m dependenciesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.finished || m.current >= len(m.dependencies) {
			m.finished = true
			return m, tea.Quit
		}

		dep := &m.dependencies[m.current]

		switch dep.Status {
		case core.Initial:
			dep.Status = core.Checking
			return m, tickAfter(500 * time.Millisecond)
		case core.Checking:
			dep.Found, dep.Status = dep.Check()
			return m, tickAfter(500 * time.Millisecond)
		case core.Installing:
			dep.Installed, dep.Status = dep.Install()
			return m, tickAfter(500 * time.Millisecond)
		case core.Done:
			m.current++
		}
	}

	// Done
	if m.current >= len(m.dependencies) {
		m.finished = true
		return m, tea.Quit
	}

	return m, nil
}

func (m dependenciesModel) View() string {
	if m.finished {
		out := "\nSummary:\n"
		for _, d := range m.dependencies {
			out += renderOutput(d) + "\n"
		}
		out += "\nPress Ctrl+C to exit.\n"

		return out
	} else {
		out := "\nChecking:\n"

		for _, d := range m.dependencies {
			out += renderOutput(d) + "\n"
		}

		return out
	}
}

func RunDependencyChecks() error {
	fmt.Println("Checking dependencies...")

	p := tea.NewProgram(newDependenciesModel())
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
		if d.Installed {
			return fmt.Sprintf("✅ %s installed.", d.Name)
		}
	case core.Checking:
		return fmt.Sprintf("🔍 Checking for %s...", d.Name)
	case core.Installing:
		return fmt.Sprintf("⬇️  Install required for %s → %s", d.Name, d.InstallURL)
	default:
		return fmt.Sprintf("ℹ  Waiting to check %s", d.Name)
	}

	return fmt.Sprintf("❗ No status update available for %s", d.Name)
}
