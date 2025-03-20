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

func computePosition(code string, transactions []models.Transaction) models.Position {
	quantity, totalCost := calculators.CalculatePosition(transactions)
	average := calculators.CalculateAveragePrice(totalCost, quantity)
	tickerData := metrics.GetTickerData(code, transactions[0].Type)
	return models.Position{
		Quantity:   quantity,
		TotalCost:  totalCost,
		Average:    average,
		Price:      tickerData.CurrentPrice,
		TotalValue: quantity * tickerData.CurrentPrice,
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
		positions[code] = computePosition(code, txs)
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

	position := computePosition(transactions[0].Code, transactions)

	c.JSON(http.StatusOK, gin.H{"position": position})
}
