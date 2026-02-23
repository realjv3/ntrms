package calc

import "math"

// Inputs holds the raw parameters needed for position sizing.
type Inputs struct {
	ATR            float64 // average true range (14-day, daily)
	Equity         float64 // account equity in dollars
	RiskPct        float64 // risk tolerance as a percentage (e.g., 2.0 = 2%)
	StopMultiplier float64 // ATR multiplier for stop distance (e.g., 1 = 1×ATR)
	EntryPrice     float64 // planned entry price per share
}

// Results holds the computed risk management variables.
type Results struct {
	Inputs
	RiskUnit     float64 // dollars risked per trade (E × p/100)
	StopDistance float64 // distance from entry to stop (x × N)
	StopPrice    float64 // entry - stop distance
	PositionSize float64 // number of shares (floor(U / stopDistance))
	TotalCost    float64 // position size × entry price
}

// Compute calculates position sizing from the given inputs.
// If StopDistance would be zero (e.g. zero ATR), PositionSize and TotalCost are 0.
func Compute(in Inputs) Results {
	riskUnit := in.Equity * (in.RiskPct / 100.0)
	stopDistance := in.StopMultiplier * in.ATR
	stopPrice := in.EntryPrice - stopDistance

	var positionSize float64

	if stopDistance > 0 {
		positionSize = math.Floor(riskUnit / stopDistance)
	}

	// cap position size to what the account can actually afford
	if in.EntryPrice > 0 {
		maxAffordable := math.Floor(in.Equity / in.EntryPrice)
		if positionSize > maxAffordable {
			positionSize = maxAffordable
		}
	}

	totalCost := positionSize * in.EntryPrice

	return Results{
		Inputs:       in,
		RiskUnit:     riskUnit,
		StopDistance: stopDistance,
		StopPrice:    stopPrice,
		PositionSize: positionSize,
		TotalCost:    totalCost,
	}
}
