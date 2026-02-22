package calc

import (
	"math"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name     string
		in       Inputs
		wantRisk float64
		wantStop float64
		wantPos  float64
		wantCost float64
	}{
		{
			name: "normal case",
			in: Inputs{
				ATR:            0.25,
				Equity:         10000,
				RiskPct:        2.0,
				StopMultiplier: 1.0,
				EntryPrice:     3.50,
			},
			wantRisk: 200.0,
			wantStop: 3.25,
			wantPos:  800.0,
			wantCost: 2800.0,
		},
		{
			name: "zero ATR edge case",
			in: Inputs{
				ATR:            0.0,
				Equity:         5000,
				RiskPct:        1.0,
				StopMultiplier: 1.0,
				EntryPrice:     2.00,
			},
			wantRisk: 50.0,
			wantStop: 2.00,
			wantPos:  0.0,
			wantCost: 0.0,
		},
		{
			name: "custom multiplier 2x",
			in: Inputs{
				ATR:            0.10,
				Equity:         25000,
				RiskPct:        2.0,
				StopMultiplier: 2.0,
				EntryPrice:     1.50,
			},
			wantRisk: 500.0,
			wantStop: 1.30,
			wantPos:  2500.0,
			wantCost: 3750.0,
		},
		{
			name: "tiny penny stock",
			in: Inputs{
				ATR:            0.02,
				Equity:         1000,
				RiskPct:        2.0,
				StopMultiplier: 1.0,
				EntryPrice:     0.05,
			},
			wantRisk: 20.0,
			wantStop: 0.03,
			wantPos:  1000.0,
			wantCost: 50.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// unit under test
			r := Compute(tt.in)

			if math.Abs(r.RiskUnit-tt.wantRisk) > 0.001 {
				t.Errorf("RiskUnit = %f, want %f", r.RiskUnit, tt.wantRisk)
			}
			if math.Abs(r.StopPrice-tt.wantStop) > 0.001 {
				t.Errorf("StopPrice = %f, want %f", r.StopPrice, tt.wantStop)
			}
			if r.PositionSize != tt.wantPos {
				t.Errorf("PositionSize = %f, want %f", r.PositionSize, tt.wantPos)
			}
			if math.Abs(r.TotalCost-tt.wantCost) > 0.001 {
				t.Errorf("TotalCost = %f, want %f", r.TotalCost, tt.wantCost)
			}
		})
	}
}
