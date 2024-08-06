package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/i9ntheory/voidsong/internal/ui"
	"github.com/i9ntheory/voidsong/pkg/generators"
)

type Model struct {
	List     list.Model
	Choice   ui.Generator
	Spinner  spinner.Model
	Quitting bool
	Err      error
}

func InitialModel() Model {
	generators := []list.Item{
		ui.Item{Generator: &generators.PrettierGenerator{}},
		ui.Item{Generator: &generators.BiomeGenerator{}},
	}

	l := ui.NewList(generators)

	return Model{
		List: l,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.List.SelectedItem().(ui.Item)
			if ok {
				m.Choice = i.Generator
				return m, tea.Batch(
					ui.RunGenerator(m.Choice),
				)
			}
		}
	case tea.WindowSizeMsg:
		h, v := ui.DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	case ui.GeneratorFinishedMsg:
		if msg.Err != nil {
			m.Err = msg.Err
			return m, tea.Quit
		}
		m.Quitting = true
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return ui.DocStyle.Render(m.List.View())
}
