package corporate_event_service

import (
	"math"

	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func calculateBasedOnType(transaction models.Transaction, corporateEvent models.CorporateEvent) (float64, float64, error) {
	var value float64
	var quantity float64
	switch corporateEvent.EventType {
	case "split":
		value = transaction.UnitPrice / corporateEvent.Ratio
		quantity = transaction.Quantity * corporateEvent.Ratio
	case "reverse_split":
		value = transaction.UnitPrice * corporateEvent.Ratio
		quantity = transaction.Quantity / corporateEvent.Ratio
	case "spin_off":
		value = float64(math.Round((transaction.UnitPrice*(1-corporateEvent.Ratio))*100) / 100)
		quantity = transaction.Quantity
	}
	return value, quantity, nil
}

func atributeTransactionValue(transaction *models.Transaction, corporateEvent models.CorporateEvent) {
	if calculatedValue, calculatedQuantity, err := calculateBasedOnType(*transaction, corporateEvent); err == nil {
		transaction.UnitPrice = calculatedValue
		transaction.Quantity = calculatedQuantity
		transaction.Adjusted = true
	}
}

func AdjustTransactionValue(corporateEvent models.CorporateEvent) []models.Transaction {
	var transactions []models.Transaction
	infrastructure.DB.Preload("Fees").
		Order("date asc").
		Where("date <= ? AND code = ?", corporateEvent.Date, corporateEvent.Code).
		Find(&transactions)

	for i := range transactions {
		atributeTransactionValue(&transactions[i], corporateEvent)
		infrastructure.DB.Save(&transactions[i])
	}

	return transactions
}
