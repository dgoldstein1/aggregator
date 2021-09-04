package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zsais/go-gin-prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Server struct {
	Coll *mongo.Collection
}

type UpdateProductRequest struct {
	Price float32 `json:"price"`
}

type UpdateProductResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// listenAndServe initializes handlers and serves RESTApi on port 8080
func (s *Server) ListenAndServe() {
	router := gin.Default()
	// add prometheus metrics ('/metrics')
	ginprometheus.NewPrometheus("gin").Use(router)
	// add handlers
	router.GET("/products/:id", s.GetProduct)
	router.PUT("/products/:id", s.UpdateProduct)
	// start router
	router.Run(":8080")
}

// GetProduct Performs an HTTP GET to retrieve the product name from an external API.
// (For this exercise the data will come from redsky.target.com, but let’s
// just pretend this is an internal resource hosted by myRetail)
func (s *Server) GetProduct(c *gin.Context) {
	stringID := c.Param("id")
	productID, err := validateIncomingProductID(stringID)
	if err != nil {
		// invalid product id: bad request
		returnErrorToClient(c, err, http.StatusBadRequest)
		return
	}
	// fetch redsky product
	redskyProduct, err := FetchRedSkyByID(productID)
	if err != nil {
		returnErrorToClient(c, err, http.StatusBadRequest)
		return
	}
	// try to find product in DB
	product, err := lookupByID(s.Coll, productID)
	if err != nil {
		returnErrorToClient(c, err, http.StatusInternalServerError)
		return
	}
	// successfully found product by ID, add name and return in response
	product.Name = redskyProduct.Product.Item.Product_Description.Title
	c.JSON(http.StatusOK, product)
}

// UpdateProduct Accepts an HTTP PUT request at the same path (/products/{id}),
// containing a JSON request body similar to the GET response, and updates the
// product’s price in the data store.
func (s *Server) UpdateProduct(c *gin.Context) {
	stringID := c.Param("id")
	productID, err := validateIncomingProductID(stringID)
	if err != nil {
		// invalid product id: bad request
		returnErrorToClient(c, err, http.StatusBadRequest)
		return
	}
	// validate incoming price
	var reqBody UpdateProductRequest
	err = c.BindJSON(&reqBody)
	if err != nil {
		returnErrorToClient(c, err, http.StatusBadRequest)
		return
	}
	// validate incoming price
	if !priceIsValid(reqBody.Price) {
		returnErrorToClient(c, errors.New("invalid price"), http.StatusBadRequest)
	}
	// update in DB
	err = updatePriceByID(s.Coll, productID, reqBody.Price)
	if err != nil {
		returnErrorToClient(c, err, http.StatusInternalServerError)
		return
	}
	// success
	c.JSON(http.StatusOK, UpdateProductResponse{fmt.Sprintf("Successfully updated price of productID %d to %f", productID, reqBody.Price)})
}

// returnErrorToClient is a helper for handlers to return error to client
func returnErrorToClient(c *gin.Context, err error, code int) {
	if err == mongo.ErrNoDocuments {
		code = http.StatusNotFound
	}
	c.JSON(code, ErrorResponse{
		Error: err.Error(),
		Code:  code,
	})
	log.Printf("returnErrorToClient: %s \n", err.Error())
}
