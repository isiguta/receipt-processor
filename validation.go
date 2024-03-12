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
		fmt.Printf("Could not validate JSON!\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Validate retailer name
	if matched, _ := regexp.MatchString("^[\\w\\s\\-]+$", r.Retailer); !matched {
		fmt.Printf("%v is invalid retailer name!\n", r.Retailer)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid retailer name"})
		c.Abort()
		return
	}

	// Validate total amount
	if matched, _ := regexp.MatchString("^\\d+\\.\\d{2}$", r.Total); !matched {
		fmt.Printf("%v is invalid total amount!\n", r.Total)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid total amount"})
		c.Abort()
		return
	}

	//Validate items
	description_compiled_regex, _ := regexp.Compile(`^[\\w\\s\\-]+$`)
	price_compiled_regex, _ := regexp.Compile(`^\\d+\\.\\d{2}$`)

	for _, item := range r.Items {
		if matched := description_compiled_regex.MatchString(item.ShortDescription); !matched {
			fmt.Printf("%v is invalid item's short description!\n", item.ShortDescription)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Short Description"})
			c.Abort()
			return
		}

		if matched := price_compiled_regex.MatchString(item.Price); !matched {
			fmt.Printf("%v is invalid item's price!\n", item.Price)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid price!"})
			c.Abort()
			return
		}
	}

	_, dateErr := time.Parse("2006-01-02", r.PurchaseDate)
	_, timeErr := time.Parse("15:04", r.PurchaseTime)

	if dateErr != nil {
		errorMessage, _ := fmt.Printf("%v is invalid purchase date!\n%v", r.PurchaseDate, dateErr.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		c.Abort()
		return
	}

	if timeErr != nil {
		errorMessage, _ := fmt.Printf("%v is invalid purchase time!\n%v", r.PurchaseTime, timeErr.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		c.Abort()
		return
	}

	c.Next()
}
