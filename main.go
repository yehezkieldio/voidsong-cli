package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Generator interface {
	Generate() (err error)
}

type AdapterRunner struct {
	adapters map[string]Generator
}

func NewAdapterRunner() *AdapterRunner {
	return &AdapterRunner{
		adapters: make(map[string]Generator),
	}
}

func (r *AdapterRunner) RegisterAdapter(name string, adapter Generator) {
	r.adapters[name] = adapter
}

func (r *AdapterRunner) Run(adapterName string) error {
	adapter, ok := r.adapters[adapterName]
	if !ok {
		return fmt.Errorf("adapter '%s' not found", adapterName)
	}
	return adapter.Generate()
}

type ESLintAdapter struct{}

type ResultMessage struct {
	msg string
}

func (e ESLintAdapter) Generate() (err error) {
	return nil
}

type Model struct {
	runner    *AdapterRunner
	adapters  []string
	cursor    int
	selected  string
	generated bool
	results   []ResultMessage
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.adapters)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.adapters[m.cursor]
			err := m.runner.Run(m.selected)
			if err != nil {
				m.results = append(m.results, ResultMessage{msg: err.Error()})
			} else {
				m.generated = true
				m.results = append(m.results, ResultMessage{msg: fmt.Sprintf("Generated %s", m.selected)})
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var s string

	s += "Select a generator to run:\n\n"

	s += "Adapters:\n"
	for i, adapter := range m.adapters {
		if i == m.cursor {
			s += fmt.Sprintf("  > %s\n", adapter)
		} else {
			s += fmt.Sprintf("    %s\n", adapter)
		}
	}

	if len(m.results) > 0 {
		s += "\n\n"
		for _, res := range m.results {
			s += fmt.Sprintf("%s\n", res.msg)
		}
	}

	return s
}

func main() {
	runner := NewAdapterRunner()
	runner.RegisterAdapter("eslint", &ESLintAdapter{})

	adapters := make([]string, 0, len(runner.adapters))
	for name := range runner.adapters {
		adapters = append(adapters, name)
	}

	model := Model{
		runner:   runner,
		adapters: adapters,
		results:  []ResultMessage{},
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
