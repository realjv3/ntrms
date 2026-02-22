package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// InputViewData carries everything the input view needs to render.
type InputViewData struct {
	Fields  []string // rendered textinput views
	ErrMsg  string   // validation or API error
	Loading bool
}

// ResultsViewData carries everything the results view needs to render.
type ResultsViewData struct {
	Ticker       string
	EntryPrice   float64
	ATR          float64
	Equity       float64
	RiskPct      float64
	StopMult     float64
	RiskUnit     float64
	StopDistance float64
	StopPrice    float64
	PositionSize float64
	TotalCost    float64
	ErrMsg       string
}

// RenderInputView renders the input form panel.
func RenderInputView(d InputViewData) string {
	title := RenderBanner()

	labels := []string{"Ticker", "Account Equity ($)", "Risk (%)", "Stop Multiplier (xATR)", "Entry Price (auto)"}

	var rows string
	for i, label := range labels {
		l := LabelStyle.Render(label)
		var field string
		if i < len(d.Fields) {
			field = d.Fields[i]
		}
		rows += fmt.Sprintf("%s  %s\n", l, field)
	}

	var errLine string
	if d.ErrMsg != "" {
		errLine = "\n" + ErrorStyle.Render(d.ErrMsg)
	}

	var status string
	if d.Loading {
		status = "\n" + ValueStyle.Render("fetching data...")
	}

	help := HelpStyle.Render("tab/shift-tab: navigate • enter: calculate • ctrl-c: quit")

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		rows,
		errLine,
		status,
		help,
	)

	return PanelStyle.Render(content)
}

// RenderResultsView renders the live dashboard panel.
func RenderResultsView(d ResultsViewData) string {
	title := RenderMiniHeader(d.Ticker)

	row := func(label, value string, style lipgloss.Style) string {
		return fmt.Sprintf("%s  %s", LabelStyle.Render(label), style.Render(value))
	}

	lines := lipgloss.JoinVertical(lipgloss.Left,
		title,
		row("Entry Price", fmt.Sprintf("$%.4f", d.EntryPrice), ValueStyle),
		row("ATR (14d)", fmt.Sprintf("%.4f", d.ATR), ValueStyle),
		row("Account Equity", fmt.Sprintf("$%.2f", d.Equity), YellowValue),
		row("Risk %", fmt.Sprintf("%.1f%%", d.RiskPct), ValueStyle),
		row("Stop Multiplier", fmt.Sprintf("%.1f×", d.StopMult), ValueStyle),
		"",
		row("Risk Unit (U)", fmt.Sprintf("$%.2f", d.RiskUnit), GreenValue),
		row("Stop Distance", fmt.Sprintf("$%.4f", d.StopDistance), RedValue),
		row("Stop Price", fmt.Sprintf("$%.4f", d.StopPrice), RedValue),
		row("Position Size", fmt.Sprintf("%.0f shares", d.PositionSize), GreenValue),
		row("Total Cost", fmt.Sprintf("$%.2f", d.TotalCost), YellowValue),
	)

	var errLine string
	if d.ErrMsg != "" {
		errLine = "\n" + ErrorStyle.Render(d.ErrMsg)
	}

	help := HelpStyle.Render("r: recalculate • q/ctrl-c: quit • price refreshes every 30s")

	content := lipgloss.JoinVertical(lipgloss.Left,
		lines,
		errLine,
		help,
	)

	return PanelStyle.Render(content)
}
