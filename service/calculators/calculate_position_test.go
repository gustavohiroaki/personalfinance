package calculators

import (
	"testing"

	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePositionTest(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 100.0, quantity)
	assert.Equal(t, 1000.0, totalCost)
}

func TestCalculatePosition(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 20.0, Operation: "C"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 200.0, quantity)
	assert.Equal(t, 3000.0, totalCost)
}

func TestCalculatePositionWithSell(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 20.0, Operation: "V"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 50.0, quantity)
	assert.Equal(t, 500.0, totalCost)
}

func TestCalculatePositionSellingAll(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 20.0, Operation: "V"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 0.0, quantity)
	assert.Equal(t, 0.0, totalCost)
}

func TestCalculatePositionSellingAllInTwoOperations(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 20.0, Operation: "V"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 20.0, Operation: "V"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 0.0, quantity)
	assert.Equal(t, 0.0, totalCost)
}

func TestCalculatePositionWithManyOperations(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 10.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 20.0, Operation: "V"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 20.0, Operation: "V"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 100, UnitPrice: 30.0, Operation: "C"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 40.0, Operation: "V"})
	transactions = append(transactions, models.Transaction{Code: "PETR4", Quantity: 50, UnitPrice: 40.0, Operation: "C"})
	quantity, totalCost := CalculatePosition(transactions)

	assert.Equal(t, 100.0, quantity)
	assert.Equal(t, 3500.0, totalCost)
}
