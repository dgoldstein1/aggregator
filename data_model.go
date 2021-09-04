package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Current_Price struct {
		Value         int    `json:"value"`
		Currency_Code string `json:"currency_code"`
	} `json:"current_price"`
}

type Server struct {
	Coll *mongo.Collection
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
