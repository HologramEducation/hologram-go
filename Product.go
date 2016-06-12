package HologramGo

import (
	"fmt"
	"os"
)

type Product map[string]interface{}
type Products map[string]interface{}

// EFFECTS: Returns product details.
func (products Products) GetProducts() {

	req := createGetRequest("/products/")

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Device{})
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}

// REQUIRES: a product id.
// EFFECTS: Returns product details.
func (product Product) GetProduct(id int) {

	req := createGetRequest("/products/" + string(id))

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Device{})
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}

// EFFECTS: Returns product categories.
func (product Product) GetProductCategories() {

	req := createGetRequest("/products/categories/")

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Device{})
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}

// EFFECTS: Returns product options.
func (product Product) GetProductOptions() {

	req := createGetRequest("/products/options/")

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Device{})
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}
