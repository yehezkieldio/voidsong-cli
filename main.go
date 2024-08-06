package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var errorTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

type item struct {
	generator Generator
}

func (i item) FilterValue() string { return i.generator.FilterValue() }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.generator.Name())

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Generator interface {
	Name() string
	Run() error
	FilterValue() string
}

type PrettierGenerator struct{}

func (g *PrettierGenerator) Name() string {
	return "Prettier"
}

func (g *PrettierGenerator) FilterValue() string {
	return "prettier"
}

func (g *PrettierGenerator) Run() error {
	fmt.Println("Finding package.json...")
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		fmt.Println(" Cannot find package.json!")
		fmt.Println(errorTextStyle.Render("\nPlease ensure you are in the root of a project with a package.json file."))
		return nil
	} else {
		fmt.Println(" Found package.json!")
	}

	fmt.Println("Finding existing configuration...")
	matches, err := filepath.Glob("*prettierrc*")
	if err != nil {
		fmt.Println(errorTextStyle.Render("\nError finding existing configuration: " + err.Error()))
		return nil
	}

	if len(matches) > 0 {
		fmt.Println(" Existing configuration found!")
		fmt.Println(" Please remove the existing configuration before running this generator.")
		fmt.Println("  - " + strings.Join(matches, "\n  - "))
	} else {
		fmt.Println(" No existing configuration found, creating new configuration...")
	}

	return nil
}

type model struct {
	list     list.Model
	choice   Generator
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() model {
	generators := []list.Item{
		item{generator: &PrettierGenerator{}},
	}

	l := list.New(generators, itemDelegate{}, 20, 14)
	l.Title = "Select a generator"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		list:    l,
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.generator
				// return m, tea.Quit
				return m, tea.Batch(
					runGenerator(m.choice),
				)
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case generatorFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
		m.quitting = true
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

type generatorFinishedMsg struct{ err error }

func runGenerator(g Generator) tea.Cmd {
	return func() tea.Msg {
		return generatorFinishedMsg{g.Run()}
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != nil {
		if err := m.choice.Run(); err != nil {
			fmt.Printf("Error running generator: %v\n", err)
			os.Exit(1)
		}
	}
}
