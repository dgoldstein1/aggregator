package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	CurrentPrice interface{} `json:"current_price"`
}

type Server struct {
	Coll *mongo.Collection
	// GetProduct    func(c *gin.Context)
	// UpdateProduct func(c *gin.Context)
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
