package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var InfoTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("213")).BorderLeft(true).BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("213")).PaddingLeft(1)

var TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("188")).BorderLeft(true).BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("188")).PaddingLeft(1)

var ErrorTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).BorderLeft(true).BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("9")).PaddingLeft(1)

var SelectionTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("213"))
var DescriptionTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
