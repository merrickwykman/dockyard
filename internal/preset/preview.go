package preset

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var exampleDimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

const (
	iconBranch = "" // Nerd Font git branch
	iconPL     = "" // Powerline right arrow
	iconNode   = "" // Node.js runtime
)

func Examples(p Preset) string {
	label := func(s string) string { return exampleDimStyle.Render("# " + s) }

	ver := "node v18"
	if !p.ShowVer {
		ver = "node"
	}

	var dirLine, cleanLine, dirtyLine, runtimeLine string

	switch p.Sep {
	case "powerline":
		sep := " " + iconPL + " "
		dirLine = "❯ ~/projects/dockyard"
		cleanLine = "❯ ~/projects/dockyard" + sep + iconBranch + " main"
		dirtyLine = "❯ ~/projects/dockyard" + sep + iconBranch + " main [+]"
		runtimeLine = "❯ ~/projects/dockyard" + sep + ver

	case "bracket":
		dirLine = "❯ [~/projects/dockyard]"
		cleanLine = "❯ [~/projects/dockyard] [" + iconBranch + " main]"
		dirtyLine = "❯ [~/projects/dockyard] [" + iconBranch + " main +]"
		runtimeLine = "❯ [~/projects/dockyard] [" + ver + "]"

	case "pure":
		dirLine = "~/projects/dockyard"
		cleanLine = "~/projects/dockyard main"
		dirtyLine = "~/projects/dockyard main*"
		runtimeLine = "~/projects/dockyard main | " + ver

	case "nofont":
		dirLine = "❯ ~/projects/dockyard"
		cleanLine = "❯ ~/projects/dockyard [git: main]"
		dirtyLine = "❯ ~/projects/dockyard [git: main*]"
		runtimeLine = "❯ ~/projects/dockyard [" + ver + "]"

	default: // plain — Nerd Font icons, default Starship style
		dirLine = "❯ ~/projects/dockyard"
		cleanLine = "❯ ~/projects/dockyard on " + iconBranch + " main"
		dirtyLine = "❯ ~/projects/dockyard on " + iconBranch + " main [+]"
		runtimeLine = "❯ ~/projects/dockyard via " + iconNode + " " + ver
	}

	lines := []string{
		label("plain directory") + "\n" + dirLine,
		label("git repo clean") + "\n" + cleanLine,
		label("git repo dirty") + "\n" + dirtyLine,
		label("language runtime") + "\n" + runtimeLine,
	}

	return strings.Join(lines, "\n")
}
