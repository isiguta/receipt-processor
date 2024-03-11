package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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
	total_pts := 0

	convertedTotal, err := strconv.ParseFloat(total, 64)

	if err != nil {

	}

	if math.Mod(convertedTotal, 1) == 0 {
		total_pts += 50
	}

	if math.Mod(convertedTotal, 0.25) == 0 {
		total_pts += 25
	}

	return total_pts
}

func processItems(items []Item) int {
	items_pts := (len(items) / 2) * 5
	for _, item := range items {
		trimmed_description := strings.TrimSpace(item.ShortDescription)
		if len(trimmed_description)%3 == 0 {
			price_value, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {

			}
			items_pts += int(math.Ceil(price_value * 0.2))
		}
	}
	return items_pts
}

func processPurchaseDateTime(date string, t string) int {
	date_pts := 0
	parsedDate, date_err := time.Parse("2006-01-02", date)
	parsedTime, time_err := time.Parse("15:04", t)

	if date_err != nil {

	}

	if time_err != nil {

	}

	if parsedDate.Day()%2 != 0 {
		date_pts += 6
	}

	if parsedTime.Hour() >= 14 && parsedTime.Hour() <= 16 {
		fmt.Printf("%v", parsedTime.Hour())
		date_pts += 10
	}

	return date_pts
}
