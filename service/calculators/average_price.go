package calculators

func CalculateAveragePrice(totalPrice float64, totalQuantity float64) float64 {
	if totalQuantity == 0 || totalPrice == 0 {
		return 0
	}
	return totalPrice / totalQuantity
}
