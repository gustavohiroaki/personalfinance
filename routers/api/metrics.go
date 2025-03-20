package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/gustavohiroaki/personalfinance/service/calculators"
	"github.com/gustavohiroaki/personalfinance/service/metrics"
	"github.com/gustavohiroaki/personalfinance/service/transaction_service"
)

func computePosition(code string, transactions []models.Transaction, currency string) models.Position {
	var getTickerResponse metrics.GetTickerResponse
	quantity, totalCost := calculators.CalculatePosition(transactions)

	if currency == "BRL" {
		getTickerResponse = metrics.GetTickerData(code + ".SA")
	} else {
		getTickerResponse = metrics.GetTickerData(code)
		currency := metrics.GetTickerData(currency + "BRL=X")
		getTickerResponse.CurrentPrice = getTickerResponse.CurrentPrice * currency.CurrentPrice
		totalCost = totalCost * currency.CurrentPrice
	}
	average := calculators.CalculateAveragePrice(totalCost, quantity)
	return models.Position{
		Quantity:   quantity,
		TotalCost:  totalCost,
		Average:    average,
		Price:      getTickerResponse.CurrentPrice,
		TotalValue: quantity * getTickerResponse.CurrentPrice,
		AssetType:  transactions[0].Type,
	}
}

func GetPosition(c *gin.Context) {
	var transactions []models.Transaction
	positions := map[string]models.Position{}

	if err := infrastructure.DB.Find(&transactions).Order("date asc").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	groupedTransactions := transaction_service.GroupTransactionsByCode(transactions)
	for code, txs := range groupedTransactions {
		currency := txs[0].Currency
		positions[code] = computePosition(code, txs, currency)
	}
	generalPositions := calculators.CalculateGeneralPosition(positions)
	c.JSON(http.StatusOK, gin.H{"positions": positions, "generalPositions": generalPositions})
}

func GetPositionByAsset(c *gin.Context) {
	var transactions []models.Transaction
	code := c.Param("id")

	if err := infrastructure.DB.Where("code = ?", code).Find(&transactions).Order("date asc").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	position := computePosition(transactions[0].Code, transactions, transactions[0].Currency)

	c.JSON(http.StatusOK, gin.H{"position": position})
}
