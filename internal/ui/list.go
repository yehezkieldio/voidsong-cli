package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

func NewList(items []list.Item) list.Model {
	l := list.New(items, GeneratorDelegate{}, 20, 14)
	l.Title = "\nWelcome to Voidsong!\nA personal configuration generator by @yehezkieldio"
	l.Styles.Title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("218"))
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	return l
}
