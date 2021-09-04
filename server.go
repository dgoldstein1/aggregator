package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"log"
	"net/http"
)

func main() {
	listenAndServe()
}

// listenAndServe initializes handlers and serves RESTApi on port 8080
func listenAndServe() {
	router := gin.Default()
	// add prometheus metrics
	ginprometheus.NewPrometheus("gin").Use(router)
	// add handlers
	router.GET("/products/:id", GetProduct)
	router.PUT("/products/:id", UpdateProduct)
	// start router
	router.Run(":8080")
}

// GetProduct Performs an HTTP GET to retrieve the product name from an external API.
// (For this exercise the data will come from redsky.target.com, but let’s
// just pretend this is an internal resource hosted by myRetail)
func GetProduct(c *gin.Context) {
	productID := c.Param("id")
	_, err := validateIncomingProductID(productID)
	if err != nil {
		// invalid product id: bad request
		returnErrorToClient(c, err, http.StatusBadRequest)
		return
	}
}

// UpdateProduct Accepts an HTTP PUT request at the same path (/products/{id}),
// containing a JSON request body similar to the GET response, and updates the
// product’s price in the data store.
func UpdateProduct(c *gin.Context) {

}

// validateProductID checks to see if incoming ID is valid
// if so, returns as int. Else returns error
func validateIncomingProductID(id string) (int, error) {
	return -1, errors.New("Not Implemented")
}

// returnErrorToClient is a helper for handlers to return error to client
func returnErrorToClient(c *gin.Context, err error, code int) {
	c.JSON(code, ErrorResponse{
		Error: err.Error(),
		Code:  code,
	})
	log.Printf("returnErrorToClient: %s \n", err.Error())
}
