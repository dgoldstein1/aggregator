
/**
 * Inserts products into mongodb
 * Usage: mongo 127.0.0.1/products docker/mongo/insert_data.js
 **/

const TOTAL_N_DOCUMENTS = 10;

/**
 * creates random number in range
 * see: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random
 * @return {int}
 **/
function getRandomInt(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

/**
 * creates random product name
 * @return {string} name
 **/
function getRandomName() {
	return (Math.random() + 1).toString(36).substring(7);
}

/**
 * Creates random product from template
 * @return {JSON} product to be inserted into mongo
 **/
function createRandomProduct() {
	return {
	    "id": getRandomInt(10000000, 99999999),
	    "current_price": {
	        "value": getRandomInt(1, 999999),
	        "currency_code": "USD"
	    }
	}
}

for (let i = 0; i < TOTAL_N_DOCUMENTS; i++) {
	// Create random document from template
	var product = createRandomProduct()
	db.products.insert(product)
	printjson(product)
}

// insert Big Lebowski
db.products.insert({
	"id" : 13860428,
	"current_price" : {
		"value" : 10,
		"currency_code" : "USD"
	}
})

print("total number of documents in DB: ", db.products.count({}))