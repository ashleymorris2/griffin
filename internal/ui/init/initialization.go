package init

import (
	"fmt"
	"github.com/ashleymorris2/booty/internal/fs"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type initModel struct {
	statuses []string
	spinner  spinner.Model
	finished bool
}

type initStep struct {
	displayName string
	run         func() error
}

type initResultMsg struct {
	Name   string
	Result string
	Err    error
}

func (m initModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		createDirCmd("Create .devsetup folder", ".devsetup"),
	)
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case initResultMsg:
		newModel := m
		newModel.statuses = append(newModel.statuses, fmt.Sprintf("%s: %s", msg.Name, msg.Result))
		//m.statuses = append(m.statuses, fmt.Sprintf("%s: %s", msg.Name, msg.Result))
		if msg.Err != nil {
			fmt.Println("Error:", msg.Err)
		}
		newModel.finished = true
		return newModel, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m initModel) View() string {
	var b strings.Builder

	if m.finished {
		b.WriteString("Setup complete!")
		b.WriteString(fmt.Sprintf("%s", strings.Join(m.statuses, "\n")))
		b.WriteString("Nice")
	} else {
		b.WriteString(fmt.Sprintf("%s Setting up...\n%s", m.spinner.View(), strings.Join(m.statuses, "\n")))
	}

	return b.String()
}

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

func newInitModel() initModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return initModel{
		spinner: s,
	}
}

func createDirCmd(name, path string) tea.Cmd {
	return func() tea.Msg {
		_, err := fs.EnsureDirExists(path)
		status := "created"
		if err != nil {
			status = "failed"
		}
		time.Sleep(5 * time.Second)
		return initResultMsg{
			Name:   name,
			Result: status,
			Err:    err,
		}
	}
}
