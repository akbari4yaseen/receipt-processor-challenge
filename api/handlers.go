package api

import (
	"net/http"

	"github.com/akbari4yaseen/receipt-processor-challenge/models"
	"github.com/akbari4yaseen/receipt-processor-challenge/services"
	"github.com/akbari4yaseen/receipt-processor-challenge/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil || receipt.Validate() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt"})
		return
	}

	id := uuid.NewString()
	points := services.CalculatePoints(receipt)
	storage.SaveReceipt(id, points)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := storage.GetPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
