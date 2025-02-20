package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func DeleteTransaction(id string) error {
	return infrastructure.DB.Delete(&models.Transaction{}, id).Error
}
