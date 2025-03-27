package corporate_event_service

import (
	"testing"

	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/gustavohiroaki/personalfinance/service/transaction_service"
	"github.com/gustavohiroaki/personalfinance/tests"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBasedOnTypeSplit(t *testing.T) {
	transaction := models.Transaction{
		UnitPrice: 10,
		Quantity:  4,
	}
	corporateEvent := models.CorporateEvent{
		EventType: "split",
		Ratio:     2,
	}
	value, quantity, err := calculateBasedOnType(transaction, corporateEvent)
	assert.Nil(t, err)
	assert.Equal(t, float64(5), value)
	assert.Equal(t, float64(8), quantity)
}

func TestCalculateBasedOnReverseSplit(t *testing.T) {
	transaction := models.Transaction{
		UnitPrice: 10,
		Quantity:  4,
	}
	corporateEvent := models.CorporateEvent{
		EventType: "reverse_split",
		Ratio:     2,
	}
	value, quantity, err := calculateBasedOnType(transaction, corporateEvent)
	assert.Nil(t, err)
	assert.Equal(t, float64(20), value)
	assert.Equal(t, float64(2), quantity)
}

func TestCalculateBasedOnSpinOff(t *testing.T) {
	transaction := models.Transaction{
		UnitPrice: 29,
		Quantity:  5,
	}
	corporateEvent := models.CorporateEvent{
		EventType: "spin_off",
		Ratio:     0.06607,
	}
	value, quantity, err := calculateBasedOnType(transaction, corporateEvent)
	assert.Nil(t, err)
	assert.Equal(t, 27.08, value)
	assert.Equal(t, float64(5), quantity)
}

func TestAdjustTransactionValue(t *testing.T) {
	tests.PrepareTests()
	gormInstance, sqlInstance := tests.PrepareDB()
	defer tests.OnClose(sqlInstance)
	transaction := models.Transaction{
		Code:      "PETR4",
		Date:      "2021-01-01",
		UnitPrice: 10,
		Quantity:  4,
		Type:      "AÇÃO",
		Operation: "C",
		Fees: models.Fees{
			Settlement: 1,
			Emolument:  1,
			Brokerage:  1,
			Iss:        1,
		},
		Currency: "BRL",
		Adjusted: false,
	}
	transactions := make([]models.Transaction, 1)
	transactions[0] = transaction
	transaction_service.CreateTransactions(transactions)
	corporateEvent := models.CorporateEvent{
		Code:      "PETR4",
		Date:      "2021-01-02",
		EventType: "split",
		Ratio:     2,
	}
	transactionsReturned := AdjustTransactionValue(corporateEvent)
	assert.Equal(t, 5.0, transactionsReturned[0].UnitPrice)
	assert.Equal(t, 8.0, transactionsReturned[0].Quantity)
	var transactionUpdated models.Transaction
	gormInstance.Where("code = ?", "PETR4").First(&transactionUpdated)

	assert.Equal(t, 5.0, transactionUpdated.UnitPrice)
	assert.Equal(t, 8.0, transactionUpdated.Quantity)
}
