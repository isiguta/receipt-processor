package main

import (
	"fmt"
	"log"
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

var receipts map[string]receipt // id:receipt

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type GetPointsByReceiptIdResponse struct {
	Points int `json:"points"`
}

func postReceipt(c *gin.Context) {
	contentLength := c.Request.ContentLength
	if contentLength == 0 {
		c.JSON(http.StatusBadRequest, "Error: body request is not present!")
		return
	}

	var incomingReceipt receipt

	if err := c.BindJSON(incomingReceipt); err != nil {
		log.Printf("Error: Invalid JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// TODO: Generate receipt id here

	// not sure we'll need it afterwards
	if id, ok := receipts[incomingReceipt.ID]; ok {
		log.Printf("%v already processed!", id) // Or what should we do?
		return
	}

	receipts[incomingReceipt.ID] = incomingReceipt

	response := ProcessReceiptResponse{
		ID: incomingReceipt.ID,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func getPointsByReceiptId(c *gin.Context) {
	contentLength := c.Request.ContentLength
	if contentLength == 0 {
		c.JSON(http.StatusBadRequest, "Error: body request is empty!")
		return
	}

	receiptId := c.Param("id")

	if receiptId == "" {
		c.JSON(http.StatusBadRequest, "Error: receiptId value is not present in the request!")
		return
	}

	r, ok := receipts[receiptId]

	if !ok {
		errorMessage := fmt.Sprintf("Error: Can't find receipt with id: %s", receiptId)
		c.JSON(http.StatusNotFound, errorMessage)
		return
	}

	points := CalculatePoints(&r)

	response := GetPointsByReceiptIdResponse{
		Points: points,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", postReceipt)
	router.GET("/receipts/:id/points", getPointsByReceiptId)
	router.Run(":8088") // in case if you have something running on 8080 already.
}
