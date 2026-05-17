package main

import (
	"fmt"
	"os"

	"github.com/MerrickWykman/dockyard/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// version is set at build time via -ldflags="-X main.version=vX.Y.Z"
var version = "dev"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(version)
		return
	}

	p := tea.NewProgram(ui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
