package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/gustavohiroaki/personalfinance/service/transaction_service"
)

func CreateBatchTransaction(c *gin.Context) {
	var transactions []models.Transaction

	if err := c.ShouldBindJSON(&transactions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	c.ShouldBindJSON(&transactions)

	if err := transaction_service.CreateTransactions(transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactions", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	transactions := make([]models.Transaction, 1)

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	transactions[0] = transaction

	if err := transaction_service.CreateTransactions(transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, transactions[0])
}

func GetTransactions(c *gin.Context) {
	queryId := c.Query("id")
	queryCode := c.Query("code")
	queryType := c.Query("type")
	queryOperation := c.Query("operation")

	transactions, err := transaction_service.GetTransactions(queryId, queryCode, queryType, queryOperation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := transaction_service.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction with ID " + id + " deleted"})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var transaction models.Transaction

	if err := infrastructure.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.ValidateTransaction(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := transaction_service.UpdateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
