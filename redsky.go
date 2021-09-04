package main

type RedSkyResponse struct {
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
