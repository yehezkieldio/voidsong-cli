package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
)

type Item struct {
	Generator Generator
}

func (i Item) FilterValue() string { return i.Generator.FilterValue() }

type GeneratorDelegate struct{}

func (d GeneratorDelegate) Height() int                             { return 4 }
func (d GeneratorDelegate) Spacing() int                            { return 0 }
func (d GeneratorDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d GeneratorDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	var (
		title, desc string
	)

	if i, ok := listItem.(Item); ok {
		title = i.Generator.Name()
		desc = i.Generator.Description()
	} else {
		return
	}

	if m.Width() <= 0 {
		return
	}

	textWidth := m.Width() - 4
	title = ansi.Truncate(title, textWidth, "…")
	var lines []string
	for i, line := range strings.Split(desc, "\n") {
		if i >= m.Height()-1 {
			break
		}
		lines = append(lines, ansi.Truncate(line, textWidth, ""))
	}
	desc = strings.Join(lines, "\n")

	var (
		isSelected = index == m.Index()
	)

	title = SelectionTextStyle.Padding(0, 0, 0, 2).Render(title)
	desc = DescriptionTextStyle.Padding(0, 0, 0, 2).Render(desc)

	if isSelected {
		title = SelectionTextStyle.Bold(true).Render(fmt.Sprintf(">%s", title))
	}

	fmt.Fprintf(w, "\n%s\n%s", title, desc)
}
