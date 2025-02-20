package corporate_event_service

import (
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
)

func CreateCorporateEvents(corporateEvents []models.CorporateEvent) error {
	for _, corporateEvent := range corporateEvents {
		if err := models.ValidateCorporateEvent(corporateEvent); err != nil {
			return err
		}
	}
	if err := infrastructure.DB.Create(&corporateEvents).Error; err != nil {
		return err
	}
	for _, corporateEvent := range corporateEvents {
		AdjustTransactionValue(corporateEvent)
	}

	return nil
}
