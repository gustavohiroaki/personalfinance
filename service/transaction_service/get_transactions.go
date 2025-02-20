package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func GetTransactions(id, queryCode, queryType, queryOperation string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := infrastructure.DB.Preload("Fees").Order("date asc").Model(&models.Transaction{})

	if id != "" {
		query = query.Where("id = ?", id)
	}
	if queryCode != "" {
		query = query.Where("code = ?", queryCode)
	}
	if queryType != "" {
		query = query.Where("type = ?", queryType)
	}
	if queryOperation != "" {
		query = query.Where("operation = ?", queryOperation)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
