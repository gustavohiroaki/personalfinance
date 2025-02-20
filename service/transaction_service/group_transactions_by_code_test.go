package transaction_service

import (
	"testing"

	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/stretchr/testify/assert"
)

func TestGroupTransactionsByCode(t *testing.T) {
	t.Run("should group transactions by code correctly", func(t *testing.T) {
		transactions := []models.Transaction{
			{Code: "AAPL", UnitPrice: 100.0},
			{Code: "GOOGL", UnitPrice: 200.0},
			{Code: "AAPL", UnitPrice: 150.0},
			{Code: "MSFT", UnitPrice: 300.0},
			{Code: "GOOGL", UnitPrice: 250.0},
		}

		result := GroupTransactionsByCode(transactions)

		assert.Len(t, result, 3)
		assert.Len(t, result["AAPL"], 2)
		assert.Len(t, result["GOOGL"], 2)
		assert.Len(t, result["MSFT"], 1)
	})

	t.Run("should return empty map when no transactions", func(t *testing.T) {
		transactions := []models.Transaction{}

		result := GroupTransactionsByCode(transactions)

		assert.Empty(t, result)
	})

	t.Run("should handle single transaction correctly", func(t *testing.T) {
		transactions := []models.Transaction{
			{Code: "AAPL", UnitPrice: 100.0},
		}

		result := GroupTransactionsByCode(transactions)

		assert.Len(t, result, 1)
		assert.Len(t, result["AAPL"], 1)
		assert.Equal(t, transactions[0], result["AAPL"][0])
	})

	t.Run("should return empty map when input is nil", func(t *testing.T) {
		var transactions []models.Transaction = nil

		result := GroupTransactionsByCode(transactions)

		assert.NotNil(t, result)
		assert.Empty(t, result)
	})
}
