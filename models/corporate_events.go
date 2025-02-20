package models

import "github.com/go-playground/validator/v10"

type CorporateEvent struct {
	DBBase
	Code            string  `json:"code" validate:"required"`
	EventType       string  `json:"event_type" validate:"required"`
	Ratio           float64 `json:"ratio"`
	AdjustmentValue float64 `json:"adjustment_value"`
	Date            string  `json:"date" validate:"required,datetime=2006-01-02"`
}

func ValidateCorporateEvent(corporateEvent CorporateEvent) error {
	validate := validator.New()
	err := validate.Struct(corporateEvent)
	if err != nil {
		return err
	}

	return err
}
