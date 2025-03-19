package transaction_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func DeleteTransaction(id string) error {
	if err := models.DeleteTransaction(infrastructure.DB, id); err != nil {
		return err
	}

	return nil
}
