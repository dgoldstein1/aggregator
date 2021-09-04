package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/patrickmn/go-cache"
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
	// cache response

	return product.Product.Item.Product_Description.Title, err
}
