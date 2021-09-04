package main

import (
	"github.com/pkg/errors"
	"strconv"
)

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Current_Price struct {
		Value         float32 `json:"value"`
		Currency_Code string  `json:"currency_code"`
	} `json:"current_price"`
}

// validateProductID checks to see if incoming ID is valid
// if so, returns as int. Else returns error
func validateIncomingProductID(id string) (int, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return -1, errors.New("invalid ID")
	}
	if i < 10000000 || i > 99999999 {
		return -1, errors.New("ID is not in valid range")
	}
	return i, nil
}

// priceIsValid is a helper that checks if price is in valid range
func priceIsValid(price float32) bool {
	return price > 0 && price <= 9999999999999999
}
