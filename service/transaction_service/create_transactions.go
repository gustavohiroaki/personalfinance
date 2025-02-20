package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func CreateTransactions(transactions []models.Transaction) error {
	for _, transaction := range transactions {
		if err := models.ValidateTransaction(transaction); err != nil {
			return err
		}
		transaction.Adjusted = false
	}
	if err := infrastructure.DB.Create(&transactions).Error; err != nil {
		return err
	}

	return nil
}
