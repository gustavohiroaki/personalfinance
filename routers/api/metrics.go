package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/gustavohiroaki/personalfinance/service/metrics"
	"github.com/gustavohiroaki/personalfinance/service/transaction_service"
)

type Position struct {
	Quantity   float64 `json:"quantity"`
	TotalCost  float64 `json:"total_cost"`
	Average    float64 `json:"average"`
	Price      float64 `json:"price"`
	TotalValue float64 `json:"total_value"`
}

func GetPosition(c *gin.Context) {
	var transactions []models.Transaction

	positions := map[string]Position{}

	if err := infrastructure.DB.Find(&transactions).Order("date asc").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	groupedTransactions := transaction_service.GroupTransactionsByCode(transactions)

	for code, transactions := range groupedTransactions {
		quantity, totalCost := metrics.CalculatePosition(transactions)
		average := metrics.CalculateAveragePrice(totalCost, quantity)
		tickerData := metrics.GetTickerData(code, transactions[0].Type)

		positions[code] = Position{
			Quantity:   quantity,
			TotalCost:  totalCost,
			Average:    average,
			Price:      tickerData.CurrentPrice,
			TotalValue: quantity * tickerData.CurrentPrice,
		}
	}

	c.JSON(http.StatusOK, gin.H{"positions": positions})
}

func GetPositionByAsset(c *gin.Context) {
	var transactions []models.Transaction
	var position Position
	code := c.Param("id")

	if err := infrastructure.DB.Where("code = ?", code).Find(&transactions).Order("date asc").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}
	quantity, totalCost := metrics.CalculatePosition(transactions)
	average := metrics.CalculateAveragePrice(totalCost, quantity)
	tickerData := metrics.GetTickerData(transactions[0].Code, transactions[0].Type)
	position = Position{
		Quantity:   quantity,
		TotalCost:  totalCost,
		Average:    average,
		Price:      tickerData.CurrentPrice,
		TotalValue: quantity * tickerData.CurrentPrice,
	}
	c.JSON(http.StatusOK, gin.H{"position": position})
}
