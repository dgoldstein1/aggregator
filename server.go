package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zsais/go-gin-prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

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
	// try to find product in DB
	product, err := lookupByID(s.Coll, productID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			returnErrorToClient(c, errors.Wrap(err, "could not find product"), http.StatusNotFound)
		} else {
			returnErrorToClient(c, errors.Wrap(err, "database error"), http.StatusInternalServerError)
		}
		return
	}
	// successfully found product by ID
	c.JSON(http.StatusOK, product)
}

// UpdateProduct Accepts an HTTP PUT request at the same path (/products/{id}),
// containing a JSON request body similar to the GET response, and updates the
// product’s price in the data store.
func (s *Server) UpdateProduct(c *gin.Context) {
	returnErrorToClient(c, errors.New("Not Implemented"), 500)
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

// returnErrorToClient is a helper for handlers to return error to client
func returnErrorToClient(c *gin.Context, err error, code int) {
	c.JSON(code, ErrorResponse{
		Error: err.Error(),
		Code:  code,
	})
	log.Printf("returnErrorToClient: %s \n", err.Error())
}
