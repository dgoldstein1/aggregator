package main

import (
	"fmt"
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Server struct {
	// product price collection
	Coll *mongo.Collection
	// map of productID : name
	RedSkyCache *cache.Cache
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
	// Create a cache with a default expiration time of 1 minute, and which
	// purges expired items every 10 minutes
	s.RedSkyCache = cache.New(1*time.Minute, 10*time.Minute)
	// initialize cache
	router := gin.Default()
	// add prometheus metrics ('/metrics')
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())
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
	redskyName, err := FetchRedSkyByID(s.RedSkyCache, productID)
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
	product.Name = redskyName
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
