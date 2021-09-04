package main

type Product struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	CurrentPrice interface{} `json:"current_price"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
