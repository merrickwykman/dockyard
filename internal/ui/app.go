package ui

import (
	"fmt"

	"github.com/MerrickWykman/dockyard/internal/detect"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appState int

const (
	stateDetecting appState = iota
	stateNotFound
	stateReady
)

var installCmd = map[string]string{
	"mac":     "brew install starship",
	"windows": "winget install starship",
	"linux":   "curl -sS https://starship.rs/install.sh | sh",
}

var headingStyle = lipgloss.NewStyle().Bold(true).MarginBottom(1)
var codeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).MarginLeft(2)
var dimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

type Model struct {
	state     appState
	detection detect.Result
	list      list.Model
}

func New() Model {
	return Model{
		state: stateDetecting,
		list:  newList(80, 24),
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg { return detect.Starship() }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case detect.Result:
		m.detection = msg
		if msg.Found {
			m.state = stateReady
		} else {
			m.state = stateNotFound
		}
		return m, nil
	}

	if m.state == stateReady {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case stateDetecting:
		return docStyle.Render("Checking for Starship…")

	case stateNotFound:
		cmd := installCmd[m.detection.Platform]
		body := fmt.Sprintf(
			"%s\n\nDockyard requires Starship to be installed.\n\nInstall it with:\n%s\n\nThen relaunch Dockyard.\n\n%s",
			headingStyle.Render("Starship not found"),
			codeStyle.Render(cmd),
			dimStyle.Render("Press Q to quit."),
		)
		return docStyle.Render(body)

	default:
		return docStyle.Render(m.list.View())
	}
}
