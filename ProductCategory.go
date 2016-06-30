package HologramGo

import (
	"fmt"
	"os"
)

type ProductCategory map[string]interface{}

type ProductCategories []interface{}

// EFFECTS: Returns product categories.
func GetProductCategories() ProductCategories {

	req := createGetRequest("/products/categories/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

///////////////////////////////////////////////////
// GENERIC PRODUCT CATEGORIES GETTER FUNCTIONS
///////////////////////////////////////////////////

func (productCategory ProductCategory) GetProductCategoryName() string {
	return productCategory["category"].(string)
}

// EFFECTS: Returns the sku of the product option.
func (productCategory ProductCategory) GetProductIdsFromCategory() []interface{} {

	return productCategory["productids"].([]interface{})
}
