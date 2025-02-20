package models

import "github.com/go-playground/validator/v10"

type Transaction struct {
	DBBase
	Code      string  `json:"code" validate:"required"`
	Type      string  `json:"type" validate:"required"`
	Operation string  `json:"operation" validate:"required,oneof=C V"`
	Date      string  `json:"date" validate:"required,datetime=2006-01-02"`
	Quantity  float64 `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
	Fees      Fees    `json:"fees" validate:"required" gorm:"foreignKey:TransactionId"`
	Adjusted  bool    `json:"adjusted"`
}

func ValidateTransaction(transaction Transaction) error {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return err
	}

	return err
}
