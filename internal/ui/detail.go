package ui

import (
	"github.com/MerrickWykman/dockyard/internal/preset"
	"github.com/charmbracelet/lipgloss"
)

var detailBorderStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderLeft(true).
	PaddingLeft(2)

var detailNameStyle = lipgloss.NewStyle().Bold(true).MarginBottom(1)

func renderDetail(p preset.Preset, width int) string {
	if p.Name == "" {
		return detailBorderStyle.Width(width).Render(dimStyle.Render("Select a preset to see details."))
	}

	body := detailNameStyle.Render(p.Name) +
		"\n" + p.Desc +
		"\n\n" + dimStyle.Render("[Enter] Apply   [Q] Quit")

	return detailBorderStyle.Width(width).Render(body)
}
