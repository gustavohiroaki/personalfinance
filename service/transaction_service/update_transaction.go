package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func UpdateTransaction(transaction models.Transaction) error {
	return infrastructure.DB.Save(&transaction).Error
}
