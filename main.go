package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/tui"
)

func main() {
	p := tea.NewProgram(&tui.InitModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error initializing tui: %v", err)
	}

}
