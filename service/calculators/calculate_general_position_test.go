package calculators

import (
	"testing"

	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculateGeneralPosition(t *testing.T) {
	positions := map[string]models.Position{
		"ABEV3": {
			Quantity:   5.0,
			TotalCost:  10.0,
			Average:    2.0,
			Price:      5.0,
			TotalValue: 25.0,
			AssetType:  "AÇÃO",
		},
		"ITUB4": {
			Quantity:   10.0,
			TotalCost:  20.0,
			Average:    2.0,
			Price:      5.0,
			TotalValue: 50.0,
			AssetType:  "AÇÃO",
		},
		"BOVA11": {
			Quantity:   5.0,
			TotalCost:  10.0,
			Average:    2.0,
			Price:      5.0,
			TotalValue: 25.0,
			AssetType:  "ETF",
		},
	}
	result := CalculateGeneralPosition(positions)
	expected := GeneralPosition{
		TotalValue:          100.0,
		TotalCost:           40.0,
		TotalGain:           60.0,
		TotalGainPercentage: 1.5,
		PositionByAssetType: map[string]PositionByAssetType{
			"AÇÃO": {
				Participation:       0.75,
				TotalValue:          75.0,
				TotalCost:           30.0,
				TotalGain:           45.0,
				TotalGainPercentage: 1.5,
			},
			"ETF": {
				Participation:       0.25,
				TotalValue:          25.0,
				TotalCost:           10.0,
				TotalGain:           15.0,
				TotalGainPercentage: 1.5,
			},
		},
	}
	assert.Equal(t, expected, result)
}
