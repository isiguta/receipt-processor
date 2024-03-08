package main

import "unicode"

func CalculatePoints(receipt *receipt) int {
	points := 0
	points += processRetailerName(receipt.Retailer)
	points += processTotalAmount(receipt.Total)
	points += processItems(receipt.Items)
	points += processPurchaseDateTime(receipt.PurchaseDate, receipt.PurchaseTime)

	return points
}

func processRetailerName(name string) int {
	retailer_points := 0
	for _, char := range name {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			retailer_points += 1
		}
	}
	return retailer_points
}

func processTotalAmount(total string) int {

	return 0
}

func processItems(items []Item) int {
	return 0
}

func processPurchaseDateTime(date string, time string) int {
	return 0
}
