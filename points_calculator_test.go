package main

import (
	"fmt"
	"testing"
)

func SetUpReceipt(retailer string, purchaseDate string, purchaseTime string, total string, items []Item) receipt {
	return receipt{
		Retailer:     retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        items,
		Total:        total,
	}
}

func TestCalculatePointsExample1(t *testing.T) {

	// Arrange
	receipt := SetUpReceipt("Target", "2024-03-11", "12:01", "35.35", []Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
	})

	// Act
	pts := CalculatePoints(&receipt)

	// Assert
	fmt.Printf("Expected: %v\n", pts)
	if pts != 28 {
		t.Fatalf("Expected and actual do not match! Expected: %v; Actual: %v", 28, pts)
	}
}

func TestCalculatePointsExample2(t *testing.T) {
	// Arrange
	receipt := SetUpReceipt("M&M Corner Market", "2023-03-20", "14:33", "9.00", []Item{
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
	})

	// Act
	pts := CalculatePoints(&receipt)

	// Assert
	fmt.Printf("Expected: %v\n", pts)
	if pts != 109 {
		t.Fatalf("Expected and actual do not match! Expected: %v; Actual: %v", 109, pts)
	}
}
