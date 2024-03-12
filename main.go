package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/receipts/process", validateReceipt, postReceipt)
	router.GET("/receipts/:id/points", getPointsByReceiptId)
	router.Run(":8088") // in case if you have something running on 8080 already.
}
