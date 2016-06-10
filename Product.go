package HologramGo

import (
	"fmt"
	"os"
)

type Product map[string]interface{}

// REQUIRES: a product id.
// EFFECTS: Returns product details.
func (product Product) GetProduct(id int) {

	req := createGetRequest("/product/" + string(id))

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

	req := createGetRequest("/product/categories/")

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
func (product Product) GetProductOptions() {

	req := createGetRequest("/product/options/")

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
