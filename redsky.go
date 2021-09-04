package main

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
)

type RedSkyProduct struct {
	Product struct {
		Item struct {
			// productID
			Tcin                string
			Product_Description struct {
				// title used in "name" field in response
				Title string `json:"title"`
			} `json:"product_description`
		} `json:"item"`
	} `json:"product"`
}

func FetchRedSkyByID(c *cache.Cache, productID int) (string, error) {
	// first check cache for product ID
	name, found := c.Get(string(productID))
	if found {
		return name.(string), nil
	}
	// if not found, reach out to API
	url := fmt.Sprintf("https://redsky.target.com/v3/pdp/tcin/%d?excludes=taxonomy,price,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics&key=candidate#_blank", productID)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad response code on redsky lookup: %d", productID)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	product := &RedSkyProduct{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		return "", err
	}
	// success: cache response
	c.Set(string(productID), product.Product.Item.Product_Description.Title, cache.DefaultExpiration)
	return product.Product.Item.Product_Description.Title, nil
}
