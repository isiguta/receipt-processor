package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func validateReceipt(c *gin.Context) {
	var r receipt
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println("Could not validate JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		c.Abort()
		return
	}

	// Validate retailer name
	if !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(r.Retailer) {
		fmt.Printf("%v is an invalid retailer name!\n", r.Retailer)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid retailer name"})
		c.Abort()
		return
	}

	// Validate total amount
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(r.Total) {
		fmt.Printf("%v is an invalid total amount!\n", r.Total)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total amount"})
		c.Abort()
		return
	}

	// Validate items
	for _, item := range r.Items {
		if !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(item.ShortDescription) {
			fmt.Printf("%v is an invalid item's short description!\n", item.ShortDescription)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid short description"})
			c.Abort()
			return
		}

		if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price) {
			fmt.Printf("%v is an invalid item's price!\n", item.Price)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			c.Abort()
			return
		}
	}

	// Validate purchase date
	if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil {
		fmt.Printf("%v is an invalid purchase date!\n%v\n", r.PurchaseDate, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase date"})
		c.Abort()
		return
	}

	// Validate purchase time
	if _, err := time.Parse("15:04", r.PurchaseTime); err != nil {
		fmt.Printf("%v is an invalid purchase time!\n%v\n", r.PurchaseTime, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase time"})
		c.Abort()
		return
	}

	c.Next()
}
