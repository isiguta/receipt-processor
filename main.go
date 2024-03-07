package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// TODO: Make a dictionary of {id:receipt} here!

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type GetReceiptPointsResponse struct {
	Points int32 `json:"points"`
}

func processReceipts(c *gin.Context) {

	response := ProcessReceiptResponse{
		ID: "123",
	}
	c.IndentedJSON(http.StatusAccepted, response)
}

func getReceiptPoints(c *gin.Context) {

	response := GetReceiptPointsResponse{
		Points: 16,
	}
	c.IndentedJSON(http.StatusAccepted, response)
}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipts)
	router.GET("/receipts/:id/points", getReceiptPoints)
	router.Run("localhost:8088") // in case if you have something running on 8080 already.
}
