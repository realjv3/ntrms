package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	bannerGreen  = lipgloss.NewStyle().Foreground(ColorGreen).Bold(true)
	bannerCyan   = lipgloss.NewStyle().Foreground(ColorCyan).Bold(true)
	bannerPurple = lipgloss.NewStyle().Foreground(ColorPurple)
)

// RenderBanner renders the full ASCII art splash banner with ninja turtle + block text.
func RenderBanner() string {
	turtle := bannerGreen.Render(strings.Join([]string{
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣤⣤⣤⣄⡀⠀⠀⠀⠀⢠⣤⣄⠀⣀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠀⣠⣴⠟⠛⠉⠁⠀⠀⠈⠉⠛⠻⣦⣄⠀⢸⡟⠙⣿⡟⣷⡀",
		"⠀⠀⠀⠀⠀⢠⣾⠏⠁⣀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡈⠻⣷⣼⣧⠀⢹⡇⣹⡇",
		"⠀⠀⠀⠀⣰⡿⠟⠛⢛⣛⣛⡿⢶⣶⣶⡶⢿⣛⣛⡛⠛⠿⢿⣿⣷⣿⣣⡿⠁",
		"⠀⠀⠀⠀⣿⠁⢀⣼⠟⣯⣝⣻⣦⣤⣤⣾⣟⣫⣭⠻⣷⡄⠈⣿⣨⣿⠋⠀⠀",
		"⠀⠀⣠⡾⠻⢷⣬⣛⣿⡿⠟⠋⠁⠀⠀⠈⠉⠛⢿⣿⣋⣵⡾⠛⢿⣅⠀⠀⠀",
		"⠀⣼⠟⠀⠀⠀⠉⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⠁⠀⠀⠀⠻⣧⠀⠀",
		"⠰⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠆⠀",
		"⠀⢻⣦⠀⠀⠀⠀⠀⢴⣤⣤⣀⣀⠀⠀⣀⣠⣤⡾⢿⡆⠀⠀⠀⠀⣴⡟⠀⠀",
		"⠀⠀⠙⢷⣤⣀⠀⠀⠀⠈⠉⠙⠛⠛⠛⠛⠉⠁⠀⠈⠁⠀⣀⣤⡾⠋⠀⠀⠀",
		"⠀⠀⠀⠀⠈⠛⠷⢶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡶⠟⠋⠁⠀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠀⠈⠛⢷⣤⣀⡀⠀⠀⢀⣠⣴⡾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠛⠛⠛⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	}, "\n"))

	text := bannerCyan.Render(strings.Join([]string{
		" ███╗   ██╗████████╗██████╗ ███╗   ███╗███████╗",
		" ████╗  ██║╚══██╔══╝██╔══██╗████╗ ████║██╔════╝",
		" ██╔██╗ ██║   ██║   ██████╔╝██╔████╔██║███████╗",
		" ██║╚██╗██║   ██║   ██╔══██╗██║╚██╔╝██║╚════██║",
		" ██║ ╚████║   ██║   ██║  ██║██║ ╚═╝ ██║███████║",
		" ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝",
	}, "\n"))

	sub := bannerPurple.Render(" ━━ Ninja Turtle Risk Management System ━━")

	rightSide := lipgloss.JoinVertical(lipgloss.Left, text, "", sub)

	return lipgloss.JoinHorizontal(lipgloss.Center, turtle, "  ", rightSide)
}

// RenderMiniHeader renders a compact colored header for the results dashboard.
func RenderMiniHeader(ticker string) string {
	face := bannerGreen.Render("⣿⣿")
	title := bannerCyan.Render("  NTRMS")
	sep := lipgloss.NewStyle().Foreground(ColorDim).Render(" ━━━ ")
	dash := bannerCyan.Render(ticker + " Dashboard")
	return face + title + sep + dash
}
