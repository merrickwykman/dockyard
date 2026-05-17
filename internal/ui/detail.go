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
var detailSectionStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("8")).MarginTop(1)
var statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

func renderDetail(p preset.Preset, width int, hasBackup bool, statusMsg string) string {
	if p.Name == "" {
		return detailBorderStyle.Width(width).Render(dimStyle.Render("Select a preset to see details."))
	}

	revertHint := dimStyle.Render("[R] Revert")
	if hasBackup {
		revertHint = "[R] Revert"
	}

	hints := "[Enter] Apply   " + revertHint + "   [Q] Quit"

	body := detailNameStyle.Render(p.Name) +
		"\n" + p.Desc +
		"\n\n" + detailSectionStyle.Render("Examples") +
		"\n" + preset.Examples(p) +
		"\n\n" + hints

	if statusMsg != "" {
		body += "\n\n" + statusStyle.Render(statusMsg)
	}

	return detailBorderStyle.Width(width).Render(body)
}
