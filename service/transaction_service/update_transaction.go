package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func UpdateTransaction(transaction models.Transaction) error {
	if err := models.UpdateTransaction(infrastructure.DB, transaction); err != nil {
		return err
	}

	return nil
}
