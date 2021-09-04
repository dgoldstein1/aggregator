package main

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

/////////////
// Product //
/////////////

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Current_Price struct {
		Value         int    `json:"value"`
		Currency_Code string `json:"currency_code"`
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
func priceIsValid(price int) bool {
	return price > 0 && price <= 9999999999999999
}

/////////////////////////
// Server and Handlers //
/////////////////////////

type Server struct {
	Coll *mongo.Collection
}

type UpdateProductRequest struct {
	Price int `json:"price"`
}

type UpdateProductResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
