package calculators

import (
	"github.com/gustavohiroaki/personalfinance/models"
)

type PositionByAssetType struct {
	Participation       float64 `json:"participation"`
	TotalValue          float64 `json:"total_value"`
	TotalCost           float64 `json:"total_cost"`
	TotalGain           float64 `json:"total_gain"`
	TotalGainPercentage float64 `json:"total_gain_percentage"`
}

type GeneralPosition struct {
	TotalValue          float64                        `json:"total_value"`
	TotalCost           float64                        `json:"total_cost"`
	TotalGain           float64                        `json:"total_gain"`
	TotalGainPercentage float64                        `json:"total_gain_percentage"`
	PositionByAssetType map[string]PositionByAssetType `json:"position_by_asset_type"`
}

func CalculateGeneralPosition(positions map[string]models.Position) GeneralPosition {
	var TotalCost float64
	var TotalValue float64
	positionByAssetType := make(map[string]PositionByAssetType)

	for _, value := range positions {
		TotalCost += value.TotalCost
		TotalValue += value.TotalValue

		if _, ok := positionByAssetType[value.AssetType]; !ok {
			positionByAssetType[value.AssetType] = PositionByAssetType{
				Participation:       0,
				TotalValue:          0,
				TotalCost:           0,
				TotalGain:           0,
				TotalGainPercentage: 0,
			}
		}

		if entry, ok := positionByAssetType[value.AssetType]; ok {
			entry.TotalCost += value.TotalCost
			entry.TotalValue += value.TotalValue
			entry.TotalGain += value.TotalValue - value.TotalCost
			positionByAssetType[value.AssetType] = entry
		}
	}

	TotalGain := TotalValue - TotalCost
	TotalGainPercentage := TotalGain / TotalCost

	for key, value := range positionByAssetType {
		value.Participation = value.TotalValue / TotalValue
		value.TotalGainPercentage = value.TotalGain / value.TotalCost
		positionByAssetType[key] = value
	}

	return GeneralPosition{
		TotalValue:          TotalValue,
		TotalCost:           TotalCost,
		TotalGain:           TotalGain,
		TotalGainPercentage: TotalGainPercentage,
		PositionByAssetType: positionByAssetType,
	}
}
