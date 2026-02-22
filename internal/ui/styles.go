package ui

import "github.com/charmbracelet/lipgloss"

// color palette
var (
	ColorGreen  = lipgloss.Color("#00FF87") // profit / position
	ColorRed    = lipgloss.Color("#FF5F56") // stop / risk
	ColorCyan   = lipgloss.Color("#00D7FF") // headers
	ColorPurple = lipgloss.Color("#AF87FF") // labels
	ColorYellow = lipgloss.Color("#FFD700") // equity / cost
	ColorDim    = lipgloss.Color("#626262") // help text
	ColorWhite  = lipgloss.Color("#FAFAFA") // default text
)

// reusable styles
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan).
			PaddingBottom(1)

	LabelStyle = lipgloss.NewStyle().
			Foreground(ColorPurple).
			Width(20)

	ValueStyle = lipgloss.NewStyle().
			Foreground(ColorWhite)

	GreenValue = lipgloss.NewStyle().
			Foreground(ColorGreen).
			Bold(true)

	RedValue = lipgloss.NewStyle().
			Foreground(ColorRed).
			Bold(true)

	YellowValue = lipgloss.NewStyle().
			Foreground(ColorYellow).
			Bold(true)

	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPurple).
			Padding(1, 2)

	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorDim).
			PaddingTop(1)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorRed).
			Bold(true)
)
