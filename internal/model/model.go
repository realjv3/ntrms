package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"ntrms/internal/api"
	"ntrms/internal/calc"
	"ntrms/internal/ui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// view states
const (
	inputView = iota
	resultsView
)

// field indices
const (
	fieldTicker = iota
	fieldEquity
	fieldRisk
	fieldStopMult
	fieldEntry
	fieldCount
)

// custom messages
type atrMsg float64
type priceMsg float64
type errMsg struct{ err error }
type tickMsg time.Time

// commands
func fetchATRCmd(apiKey, ticker string) tea.Cmd {
	return func() tea.Msg {
		atr, err := api.FetchATR(apiKey, ticker)
		if err != nil {
			return errMsg{err}
		}

		return atrMsg(atr)
	}
}

func fetchPriceCmd(apiKey, ticker string) tea.Cmd {
	return func() tea.Msg {
		price, err := api.FetchPrice(apiKey, ticker)
		if err != nil {
			return errMsg{err}
		}

		return priceMsg(price)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(30*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Model is the root Bubbletea model.
type Model struct {
	apiKey string
	view   int
	focus  int
	inputs [fieldCount]textinput.Model
	err    string

	// parsed form values
	ticker  string
	equity  float64
	riskPct float64
	stopMul float64
	entry   float64 // 0 means auto from API

	// api state
	atr      float64
	price    float64
	gotATR   bool
	gotPrice bool
	loading  bool

	// computed
	results calc.Results
}

// New creates a new model with defaults.
func New(apiKey string) Model {
	m := Model{apiKey: apiKey}

	for i := range m.inputs {
		t := textinput.New()
		t.CharLimit = 20

		switch i {
		case fieldTicker:
			t.Placeholder = "e.g. AAPL"
			t.Prompt = ""
		case fieldEquity:
			t.Placeholder = "10000"
			t.Prompt = ""
		case fieldRisk:
			t.Placeholder = "2"
			t.Prompt = ""
			t.SetValue("2")
		case fieldStopMult:
			t.Placeholder = "1"
			t.Prompt = ""
			t.SetValue("1")
		case fieldEntry:
			t.Placeholder = "auto"
			t.Prompt = ""
		}
		m.inputs[i] = t
	}

	m.inputs[fieldTicker].Focus()
	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.view {
	case inputView:
		return m.updateInput(msg)
	case resultsView:
		return m.updateResults(msg)
	}
	return m, nil
}

func (m Model) View() string {
	switch m.view {
	case resultsView:
		return m.viewResults()
	default:
		return m.viewInput()
	}
}

// --- input view ---

func (m Model) updateInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "down":
			m.focus = (m.focus + 1) % fieldCount
			return m, m.syncFocus()
		case "shift+tab", "up":
			m.focus = (m.focus - 1 + fieldCount) % fieldCount
			return m, m.syncFocus()
		case "enter":
			return m.submit()
		}

	case atrMsg:
		m.atr = float64(msg)
		m.gotATR = true
		return m.tryCompute()

	case priceMsg:
		m.price = float64(msg)
		m.gotPrice = true
		return m.tryCompute()

	case errMsg:
		m.err = msg.err.Error()
		m.loading = false
		return m, nil
	}

	// forward to focused input
	cmd := m.updateFocusedInput(msg)
	return m, cmd
}

func (m *Model) syncFocus() tea.Cmd {
	var cmds []tea.Cmd
	for i := range m.inputs {
		if i == m.focus {
			cmds = append(cmds, m.inputs[i].Focus())
		} else {
			m.inputs[i].Blur()
		}
	}
	return tea.Batch(cmds...)
}

func (m *Model) updateFocusedInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.inputs[m.focus], cmd = m.inputs[m.focus].Update(msg)
	return cmd
}

func (m Model) submit() (tea.Model, tea.Cmd) {
	// validate ticker
	ticker := strings.TrimSpace(strings.ToUpper(m.inputs[fieldTicker].Value()))
	if ticker == "" {
		m.err = "ticker is required"

		return m, nil
	}
	m.ticker = ticker

	// validate equity
	eq, err := strconv.ParseFloat(strings.TrimSpace(m.inputs[fieldEquity].Value()), 64)
	if err != nil || eq <= 0 {
		m.err = "equity must be a positive number"

		return m, nil
	}
	m.equity = eq

	// validate risk %
	rp, err := strconv.ParseFloat(strings.TrimSpace(m.inputs[fieldRisk].Value()), 64)
	if err != nil || rp <= 0 || rp > 100 {
		m.err = "risk % must be between 0 and 100"

		return m, nil
	}
	m.riskPct = rp

	// validate stop multiplier
	sm, err := strconv.ParseFloat(strings.TrimSpace(m.inputs[fieldStopMult].Value()), 64)
	if err != nil || sm <= 0 {
		m.err = "stop multiplier must be a positive number"

		return m, nil
	}
	m.stopMul = sm

	// validate optional entry price
	entryStr := strings.TrimSpace(m.inputs[fieldEntry].Value())
	if entryStr == "" || strings.EqualFold(entryStr, "auto") {
		m.entry = 0 // auto
	} else {
		ep, err := strconv.ParseFloat(entryStr, 64)
		if err != nil || ep <= 0 {
			m.err = "entry price must be a positive number or blank for auto"

			return m, nil
		}
		m.entry = ep
	}

	m.err = ""
	m.loading = true
	m.gotATR = false
	m.gotPrice = false

	return m, tea.Batch(
		fetchATRCmd(m.apiKey, m.ticker),
		fetchPriceCmd(m.apiKey, m.ticker),
	)
}

func (m Model) tryCompute() (tea.Model, tea.Cmd) {
	if !m.gotATR || !m.gotPrice {
		return m, nil
	}

	entryPrice := m.price
	if m.entry > 0 {
		entryPrice = m.entry
	}

	m.results = calc.Compute(calc.Inputs{
		ATR:            m.atr,
		Equity:         m.equity,
		RiskPct:        m.riskPct,
		StopMultiplier: m.stopMul,
		EntryPrice:     entryPrice,
	})

	m.loading = false
	m.view = resultsView
	return m, tickCmd()
}

func (m Model) viewInput() string {
	fields := make([]string, fieldCount)

	for i := range m.inputs {
		fields[i] = m.inputs[i].View()
	}

	return ui.RenderInputView(ui.InputViewData{
		Fields:  fields,
		ErrMsg:  m.err,
		Loading: m.loading,
	})
}

// --- results view ---

func (m Model) updateResults(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			m.view = inputView
			m.gotATR = false
			m.gotPrice = false

			return m, m.syncFocus()
		}

	case priceMsg:
		m.price = float64(msg)
		entryPrice := m.price
		if m.entry > 0 {
			entryPrice = m.entry
		}
		m.results = calc.Compute(calc.Inputs{
			ATR:            m.atr,
			Equity:         m.equity,
			RiskPct:        m.riskPct,
			StopMultiplier: m.stopMul,
			EntryPrice:     entryPrice,
		})
		m.err = ""

		return m, nil

	case errMsg:
		m.err = fmt.Sprintf("refresh error: %v", msg.err)

		return m, nil

	case tickMsg:
		return m, tea.Batch(
			fetchPriceCmd(m.apiKey, m.ticker),
			tickCmd(),
		)
	}

	return m, nil
}

func (m Model) viewResults() string {
	r := m.results
	return ui.RenderResultsView(ui.ResultsViewData{
		Ticker:       m.ticker,
		EntryPrice:   r.EntryPrice,
		ATR:          r.ATR,
		Equity:       r.Equity,
		RiskPct:      r.RiskPct,
		StopMult:     r.StopMultiplier,
		RiskUnit:     r.RiskUnit,
		StopDistance: r.StopDistance,
		StopPrice:    r.StopPrice,
		PositionSize: r.PositionSize,
		TotalCost:    r.TotalCost,
		ErrMsg:       m.err,
	})
}
