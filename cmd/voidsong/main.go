package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/i9ntheory/voidsong/internal/app"
)

func main() {
	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m, ok := m.(app.Model); ok && m.Choice != nil {
		if err := m.Choice.Run(); err != nil {
			fmt.Printf("Error running generator: %v\n", err)
			os.Exit(1)
		}
	}
}
