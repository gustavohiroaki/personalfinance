package calculators

import "github.com/gustavohiroaki/personalfinance/models"

func CalculatePosition(transactions []models.Transaction) (float64, float64) {
	var quantity float64
	var totalCost float64
	for _, transaction := range transactions {
		if transaction.Operation == "C" {
			quantity += transaction.Quantity
			totalCost += transaction.UnitPrice * float64(transaction.Quantity)
		}
		if transaction.Operation == "V" {
			average := CalculateAveragePrice(totalCost, quantity)
			quantity -= transaction.Quantity
			totalCost -= average * transaction.Quantity

			if quantity <= 0 {
				quantity = 0
				totalCost = 0
			}
		}
	}

	if quantity == 0 {
		return 0.0, 0.0
	}
	return quantity, totalCost
}
