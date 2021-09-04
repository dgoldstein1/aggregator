package main

import (
	"encoding/json"
	"fmt"
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

func FetchRedSkyByID(productID int) (*RedSkyProduct, error) {
	url := fmt.Sprintf("https://redsky.target.com/v3/pdp/tcin/%d?excludes=taxonomy,price,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics&key=candidate#_blank", productID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad response code on redsky lookup: %d", productID)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	product := &RedSkyProduct{}
	err = json.Unmarshal(body, &product)
	return product, err
}
