package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

func NewList(items []list.Item) list.Model {
	l := list.New(items, ItemDelegate{}, 20, 14)
	l.Title = "Select a generator"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	return l
}
