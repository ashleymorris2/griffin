package ui

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/core"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type stepStage int

const (
	stageInit stepStage = iota
	stageChecking
	stageInstalling
	stageDone
)

// Represents a single step in the init process
type initStep struct {
	name       string
	command    string
	status     stepStage
	installURL string
}

type initModel struct {
	steps    []initStep
	current  int
	statuses []string
	spinner  spinner.Model
}

type tickMsg struct{}

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return initModel{
		steps: []initStep{
			{name: "WSL", command: "wsl"},
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
	switch msg := msg.(type) {

	case tickMsg:
		if m.current >= len(m.steps) {
			return m, tea.Quit
		}

		currentStep := &m.steps[m.current]

		switch currentStep.status {
		case stageInit:
			currentStep.status = stageChecking
			return m, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
				return tickMsg{}
			})
		case stageChecking:
			currentStep.found = core.DependencyExists(step.command)

			if core.DependencyExists(currentStep.name) {
				m.statuses = append(m.statuses, fmt.Sprintf("%s found", currentStep.name))
			} else {

			}

			m.current++

			// Schedule next tick after short delay
			return m, tea.Tick(time.Millisecond*5000, func(t time.Time) tea.Msg {
				return tickMsg{}
			})
		case stageInstalling:
		case stageDone:
		}

	default:
		// Update spinner
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m initModel) View() string {
	if m.done {
		out := "\nSummary:\n"
		for _, s := range m.statuses {
			out += s + "\n"
		}
		out += "\nPress Ctrl+C to exit.\n"
		return out
	}

	if m.current >= len(m.steps) {
		return fmt.Sprintf("%s Finishing up.", m.spinner.View())
	}

	currentStep := m.steps[m.current].name
	return fmt.Sprintf("%s Checking for %s...\n", m.spinner.View(), currentStep)
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
