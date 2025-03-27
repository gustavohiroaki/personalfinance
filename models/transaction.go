package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Transaction struct {
	DBBase
	Code      string  `json:"code" validate:"required"`
	Type      string  `json:"type" validate:"required"`
	Operation string  `json:"operation" validate:"required,oneof=C V"`
	Date      string  `json:"date" validate:"required,datetime=2006-01-02"`
	Quantity  float64 `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
	Fees      Fees    `json:"fees" validate:"required" gorm:"foreignKey:TransactionId"`
	Currency  string  `json:"currency" validate:"required"`
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

func CreateTransaction(DB *gorm.DB, transaction *Transaction) error {
	err := ValidateTransaction(*transaction)
	if err != nil {
		return err
	}

	return DB.Create(transaction).Error
}

func CreateTransactions(DB *gorm.DB, transactions *[]Transaction) error {
	for _, transaction := range *transactions {
		err := ValidateTransaction(transaction)
		if err != nil {
			return err
		}
	}

	return DB.Create(transactions).Error
}

func FindTransactions(DB *gorm.DB, id, queryCode, queryType, queryOperation string) ([]Transaction, error) {
	var transactions []Transaction
	query := DB.Preload("Fees").Order("date asc").Model(&Transaction{})

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

func DeleteTransaction(DB *gorm.DB, id string) error {
	return DB.Delete(&Transaction{}, id).Error
}

func UpdateTransaction(DB *gorm.DB, transaction Transaction) error {
	err := ValidateTransaction(transaction)
	if err != nil {
		return err
	}

	return DB.Save(&transaction).Error
}
