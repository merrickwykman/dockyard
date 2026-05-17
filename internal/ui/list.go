package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type preset struct {
	name string
	desc string
}

func (p preset) Title() string       { return p.name }
func (p preset) Description() string { return p.desc }
func (p preset) FilterValue() string { return p.name }

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func newList(width, height int) list.Model {
	items := []list.Item{
		preset{"Nerd Font Symbols", "Icons and glyphs from Nerd Fonts"},
		preset{"Pastel Powerline", "Soft colors with Powerline separators"},
		preset{"Tokyo Night", "Dark theme inspired by the city at night"},
		preset{"Gruvbox Rainbow", "Warm retro palette with rainbow segments"},
		preset{"Pure", "Minimal single-line prompt"},
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Dockyard — Starship Presets"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	return l
}
