package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func GetTransactions(id, queryCode, queryType, queryOperation string) ([]models.Transaction, error) {
	transactions, err := models.FindTransactions(infrastructure.DB, id, queryCode, queryType, queryOperation)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
