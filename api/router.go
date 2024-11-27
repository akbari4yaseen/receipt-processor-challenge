package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/receipts/process", ProcessReceipt)
	router.GET("/receipts/:id/points", GetPoints)

	return router
}
