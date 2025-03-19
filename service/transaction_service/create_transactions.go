package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func CreateTransactions(transactions []models.Transaction) error {
	if err := models.CreateTransactions(infrastructure.DB, &transactions); err != nil {
		return err
	}

	return nil
}
