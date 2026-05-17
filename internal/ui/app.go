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
	stateFontChecking
	stateFontWarning
	stateFontInstalling
	stateReady
)

type fontInstallResult struct{ err error }

var starshipInstallCmd = map[string]string{
	"mac":     "brew install starship",
	"windows": "winget install starship",
	"linux":   "curl -sS https://starship.rs/install.sh | sh",
}

var headingStyle = lipgloss.NewStyle().Bold(true).MarginBottom(1)
var codeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).MarginLeft(2)
var dimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

type Model struct {
	state      appState
	starship   detect.Result
	fontResult detect.FontResult
	installErr string
	list       list.Model
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
		case "c", "enter":
			if m.state == stateFontWarning {
				m.state = stateReady
				return m, nil
			}
		case "i":
			if m.state == stateFontWarning && m.fontResult.Platform != "windows" {
				m.state = stateFontInstalling
				platform := m.fontResult.Platform
				return m, func() tea.Msg {
					return fontInstallResult{err: detect.InstallFont(platform)}
				}
			}
		}

	case detect.Result:
		m.starship = msg
		if !msg.Found {
			m.state = stateNotFound
			return m, nil
		}
		m.state = stateFontChecking
		return m, func() tea.Msg { return detect.Font() }

	case detect.FontResult:
		m.fontResult = msg
		if msg.Supported {
			m.state = stateReady
		} else {
			m.state = stateFontWarning
		}
		return m, nil

	case fontInstallResult:
		if msg.err != nil {
			m.installErr = msg.err.Error()
			m.state = stateFontWarning
			return m, nil
		}
		m.state = stateFontChecking
		return m, func() tea.Msg { return detect.Font() }
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
		cmd := starshipInstallCmd[m.starship.Platform]
		body := fmt.Sprintf(
			"%s\n\nDockyard requires Starship to be installed.\n\nInstall it with:\n%s\n\nThen relaunch Dockyard.\n\n%s",
			headingStyle.Render("Starship not found"),
			codeStyle.Render(cmd),
			dimStyle.Render("Press Q to quit."),
		)
		return docStyle.Render(body)

	case stateFontChecking:
		return docStyle.Render("Checking for Nerd Font…")

	case stateFontInstalling:
		return docStyle.Render("Installing font…")

	case stateFontWarning:
		return docStyle.Render(m.fontWarningView())

	default:
		return docStyle.Render(m.list.View())
	}
}

func (m Model) fontWarningView() string {
	body := headingStyle.Render("Nerd Font not found") +
		"\n\nDockyard works best with a Nerd Font installed.\n" +
		"Without one, icons in the preset list may appear as broken characters.\n\n"

	switch m.fontResult.Platform {
	case "mac":
		body += "Install automatically:\n" + codeStyle.Render("brew install --cask font-jetbrains-mono-nerd-font") + "\n\n"
	case "linux":
		body += "Install automatically (downloads JetBrainsMono Nerd Font).\n\n"
	case "windows":
		body += "Download a Nerd Font from " + codeStyle.Render("https://www.nerdfonts.com/font-downloads") +
			"\nExtract the .ttf files and install by right-clicking → Install for all users.\n\n"
	}

	if m.installErr != "" {
		body += dimStyle.Render("Install failed: "+m.installErr) + "\n\n"
	}

	if m.fontResult.Platform != "windows" {
		body += "[I] Install   [C] Continue anyway   [Q] Quit"
	} else {
		body += "[C] Continue anyway   [Q] Quit"
	}

	return body
}
