package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Generator interface {
	Name() string
	Run() error
	FilterValue() string
}

type GeneratorFinishedMsg struct{ Err error }

func RunGenerator(g Generator) tea.Cmd {
	return func() tea.Msg {
		return GeneratorFinishedMsg{g.Run()}
	}
}
