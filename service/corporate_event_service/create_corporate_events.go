package corporate_event_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func CreateCorporateEvents(corporateEvents []models.CorporateEvent) error {
	if err := models.CreateCorporateEvents(infrastructure.DB, &corporateEvents); err != nil {
		return err
	}
	for _, corporateEvent := range corporateEvents {
		AdjustTransactionValue(corporateEvent)
	}

	return nil
}
