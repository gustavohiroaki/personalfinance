package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/models"
	"github.com/gustavohiroaki/personalfinance/service/corporate_event_service"
)

func CreateBatchCorporateEvent(c *gin.Context) {
	var corporateEvents []models.CorporateEvent

	if err := c.ShouldBindJSON(&corporateEvents); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := corporate_event_service.CreateCorporateEvents(corporateEvents); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create corporate event"})
		return
	}

	c.JSON(http.StatusOK, corporateEvents)
}

func CreateCorporateEvent(c *gin.Context) {
	var corporateEvent models.CorporateEvent

	if err := c.ShouldBindJSON(&corporateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.ValidateCorporateEvent(corporateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	corporate_event_service.AdjustTransactionValue(corporateEvent)

	if err := infrastructure.DB.Create(&corporateEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create corporate event"})
		return
	}

	c.JSON(http.StatusOK, corporateEvent)
}

func GetCorporateEvents(c *gin.Context) {
	var corporateEvent []models.CorporateEvent

	queryEventType := c.Query("event_type")
	queryCode := c.Query("code")

	query := infrastructure.DB.Order("date asc").Model(&models.CorporateEvent{})

	if queryEventType != "" {
		query = query.Where("event_type = ?", queryEventType)
	}
	if queryCode != "" {
		query = query.Where("code = ?", queryCode)
	}

	if err := query.Find(&corporateEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get corporate event"})
		return
	}
	c.JSON(http.StatusOK, corporateEvent)
}

func DeleteCorporateEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	infrastructure.DB.Delete(&models.CorporateEvent{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Corporate event with ID " + id + " deleted"})
}

func UpdateCorporateEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var corporateEvent models.CorporateEvent

	if err := infrastructure.DB.First(&corporateEvent, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Corporate event not found"})
		return
	}

	if err := c.ShouldBindJSON(&corporateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.ValidateCorporateEvent(corporateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := infrastructure.DB.Save(&corporateEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update corporate event"})
		return
	}

	c.JSON(http.StatusOK, corporateEvent)
}
