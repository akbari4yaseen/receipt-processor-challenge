package services

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/akbari4yaseen/receipt-processor-challenge/models"
)

// CalculatePoints applies the points calculation rules to a receipt
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	reg := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(reg.FindAllString(receipt.Retailer, -1))

	// Rule 2: 50 points if total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 1.0) == 0 {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions with trimmed length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if purchase time is after 2:00 PM and before 4:00 PM
	hour, _ := strconv.Atoi(strings.Split(receipt.PurchaseTime, ":")[0])
	if hour == 14 || hour == 15 {
		points += 10
	}

	return points
}
