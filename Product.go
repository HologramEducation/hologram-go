package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

// Product implements the Product type returned in the response.
type Product map[string]interface{}

// Products is just a list of Product(s).
type Products []interface{}

// GetProducts returns product details.
func GetProducts() Products {

	req := createGetRequest("/products/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

// GetProduct returns product details response.
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

// GetProductId returns the id of the product.
func (product Product) GetProductId() float64 {
	return product["id"].(float64)
}

// GetProductSku returns the sku of the product.
func (product Product) GetProductSku() string {
	return product["sku"].(string)
}

// GetProductName returns the name of the product.
func (product Product) GetProductName() string {
	return product["name"].(string)
}

// GetProductDescription returns the description of the product.
func (product Product) GetProductDescription() string {
	return product["description"].(string)
}

// GetProductPrice returns the price of the product.
func (product Product) GetProductPrice() string {
	return product["price"].(string)
}

// GetProductImageUrl returns the sku of the product.
func (product Product) GetProductImageUrl() string {
	return product["imageurl"].(string)
}

// GetProductInvoiceDescription returns the invoice description of the product.
func (product Product) GetProductInvoiceDescription() string {
	return product["invoice_description"].(string)
}

// GetProductPreorderDetails returns the invoice description of the product.
func (product Product) GetProductPreorderDetails() string {
	return product["preorder_details"].(string)
}
