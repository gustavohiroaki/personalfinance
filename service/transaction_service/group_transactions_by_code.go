package transaction_service

import "github.com/gustavohiroaki/personalfinance/models"

func GroupTransactionsByCode(transactions []models.Transaction) map[string][]models.Transaction {
	if transactions == nil {
		return make(map[string][]models.Transaction)
	}

	estimatedCapacity := len(transactions) / 2
	groupedByCode := make(map[string][]models.Transaction, estimatedCapacity)

	for _, tx := range transactions {
		groupedByCode[tx.Code] = append(groupedByCode[tx.Code], tx)
	}

	return groupedByCode
}
