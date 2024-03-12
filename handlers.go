package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]receipt) // id:receipt

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

	validatedReceipt, _ := c.Get("validatedReceipt")
	incomingReceipt := validatedReceipt.(receipt)

	// Generate receipt id here
	incomingReceipt.ID = uuid.New().String()

	receipts[incomingReceipt.ID] = incomingReceipt
	response := ProcessReceiptResponse{
		ID: incomingReceipt.ID,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func getPointsByReceiptId(c *gin.Context) {
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
