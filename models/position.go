package models

type Position struct {
	Quantity   float64 `json:"quantity"`
	TotalCost  float64 `json:"total_cost"`
	Average    float64 `json:"average"`
	Price      float64 `json:"price"`
	TotalValue float64 `json:"total_value"`
	AssetType  string  `json:"asset_type"`
}
