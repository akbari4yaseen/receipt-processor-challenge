package models

import (
	"errors"
	"regexp"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Validate checks the validity of a Receipt object
func (r *Receipt) Validate() error {
	if r.Retailer == "" || r.PurchaseDate == "" || r.PurchaseTime == "" || r.Total == "" || len(r.Items) == 0 {
		return errors.New("missing required fields")
	}

	// Validate total format
	match, _ := regexp.MatchString(`^\d+\.\d{2}$`, r.Total)
	if !match {
		return errors.New("invalid total format")
	}

	// Validate items
	for _, item := range r.Items {
		if item.ShortDescription == "" || item.Price == "" {
			return errors.New("item has missing fields")
		}
		match, _ := regexp.MatchString(`^\d+\.\d{2}$`, item.Price)
		if !match {
			return errors.New("invalid item price format")
		}
	}
	return nil
}
