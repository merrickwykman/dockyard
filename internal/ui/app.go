package ui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MerrickWykman/dockyard/internal/config"
	"github.com/MerrickWykman/dockyard/internal/detect"
	"github.com/MerrickWykman/dockyard/internal/preset"
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
	stateLoading
	stateReady
)

type fontInstallResult struct{ err error }
type presetsLoaded struct{ items []preset.Preset }
type applyResult struct{ err error }
type revertResult struct{ err error }

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
	width      int
	height     int
	starship   detect.Result
	fontResult detect.FontResult
	installErr string
	list       list.Model
	hasBackup  bool
	statusMsg  string
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

func checkBackup() bool {
	p, err := config.Path()
	if err != nil {
		return false
	}
	_, err = os.Stat(p + ".bak")
	return err == nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width/2-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "c":
			if m.state == stateFontWarning {
				m.state = stateLoading
				return m, func() tea.Msg { return presetsLoaded{preset.List()} }
			}

		case "enter":
			if m.state == stateFontWarning {
				m.state = stateLoading
				return m, func() tea.Msg { return presetsLoaded{preset.List()} }
			}
			if m.state == stateReady {
				sel, ok := m.list.SelectedItem().(presetItem)
				if !ok {
					break
				}
				name := sel.Preset.Name
				return m, func() tea.Msg {
					out, err := exec.Command("starship", "preset", name).Output()
					if err != nil {
						return applyResult{err}
					}
					return applyResult{err: config.Write(string(out))}
				}
			}

		case "i":
			if m.state == stateFontWarning && m.fontResult.Platform != "windows" {
				m.state = stateFontInstalling
				platform := m.fontResult.Platform
				return m, func() tea.Msg {
					return fontInstallResult{err: detect.InstallFont(platform)}
				}
			}

		case "r":
			if m.state == stateReady && m.hasBackup {
				return m, func() tea.Msg {
					return revertResult{err: config.Revert()}
				}
			}

		case "up", "down", "k", "j":
			if m.state == stateReady {
				m.statusMsg = ""
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
			m.state = stateLoading
			return m, func() tea.Msg { return presetsLoaded{preset.List()} }
		}
		m.state = stateFontWarning
		return m, nil

	case fontInstallResult:
		if msg.err != nil {
			m.installErr = msg.err.Error()
			m.state = stateFontWarning
			return m, nil
		}
		m.state = stateFontChecking
		return m, func() tea.Msg { return detect.Font() }

	case presetsLoaded:
		m.list.SetItems(toListItems(msg.items))
		m.hasBackup = checkBackup()
		m.state = stateReady
		return m, nil

	case applyResult:
		if msg.err != nil {
			m.statusMsg = "Error: " + msg.err.Error()
		} else {
			m.statusMsg = "✓ Preset applied. Restart your terminal to see changes."
			m.hasBackup = true
		}
		return m, nil

	case revertResult:
		if msg.err != nil {
			m.statusMsg = "Error: " + msg.err.Error()
		} else {
			m.statusMsg = "✓ Config reverted."
			m.hasBackup = checkBackup()
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
	case stateLoading:
		return docStyle.Render("Loading presets…")
	default:
		return m.readyView()
	}
}

func (m Model) readyView() string {
	half := m.width / 2
	listPane := lipgloss.NewStyle().Width(half).Render(m.list.View())
	sel, _ := m.list.SelectedItem().(presetItem)
	detailPane := renderDetail(sel.Preset, m.width-half, m.hasBackup, m.statusMsg)
	return lipgloss.JoinHorizontal(lipgloss.Top, listPane, detailPane)
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
