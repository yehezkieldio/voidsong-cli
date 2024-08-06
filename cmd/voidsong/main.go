package voidsong

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/i9ntheory/voidsong/internal/app"
)

func Execute() error {
	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return errors.New("voidsong: error running program")
	}

	if m, ok := m.(app.Model); ok && m.Choice != nil {
		if err := m.Choice.Run(); err != nil {
			fmt.Printf("Error running generator: %v\n", err)
			return errors.New("voidsong: error running generator")
		}
	}

	return nil
}
