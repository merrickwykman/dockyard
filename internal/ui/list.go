package ui

import (
	"github.com/MerrickWykman/dockyard/internal/preset"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type presetItem struct{ Preset preset.Preset }

func (p presetItem) Title() string       { return p.Preset.Name }
func (p presetItem) Description() string { return p.Preset.Desc }
func (p presetItem) FilterValue() string { return p.Preset.Name }

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func newList(width, height int) list.Model {
	l := list.New(nil, list.NewDefaultDelegate(), width, height)
	l.Title = "Dockyard — Starship Presets"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	return l
}

func toListItems(presets []preset.Preset) []list.Item {
	items := make([]list.Item, len(presets))
	for i, p := range presets {
		items[i] = presetItem{p}
	}
	return items
}
