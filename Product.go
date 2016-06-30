package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

type Product map[string]interface{}

type Products []interface{}

// EFFECTS: Returns product details.
func GetProducts() Products {

	req := createGetRequest("/products/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

// REQUIRES: a product id.
// EFFECTS: Returns product details response.
func GetProduct(id int) Product {

	req := createGetRequest("/products/" + strconv.Itoa(id))

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

///////////////////////////////////////////////////
// GENERIC PRODUCT GETTER FUNCTIONS
///////////////////////////////////////////////////

// EFFECTS: Returns the id of the product.
func (product Product) GetProductId() float64 {
	return product["id"].(float64)
}

// EFFECTS: Returns the sku of the product.
func (product Product) GetProductSku() string {
	return product["sku"].(string)
}

// EFFECTS: Returns the name of the product.
func (product Product) GetProductName() string {
	return product["name"].(string)
}

// EFFECTS: Returns the description of the product.
func (product Product) GetProductDescription() string {
	return product["description"].(string)
}

// EFFECTS: Returns the price of the product.
func (product Product) GetProductPrice() string {
	return product["price"].(string)
}

// EFFECTS: Returns the sku of the product.
func (product Product) GetProductImageUrl() string {
	return product["imageurl"].(string)
}

// EFFECTS: Returns the invoice description of the product.
func (product Product) GetProductInvoiceDescription() string {
	return product["invoice_description"].(string)
}

// EFFECTS: Returns the invoice description of the product.
func (product Product) GetProductPreorderDetails() string {
	return product["preorder_details"].(string)
}
