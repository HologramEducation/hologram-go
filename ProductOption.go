package HologramGo

import (
	"fmt"
	"os"
)

type ProductOption map[string]interface{}
type ProductOptions []interface{}

// EFFECTS: Returns product options.
func GetProductOptions() ProductOptions {

	req := createGetRequest("/products/options/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

///////////////////////////////////////////////////
// PRODUCT OPTION GETTER FUNCTIONS
///////////////////////////////////////////////////

func (productOption ProductOption) GetProductIdFromOption() float64 {
	return productOption["productid"].(float64)
}

// EFFECTS: Returns the sku of the product option.
func (productOption ProductOption) GetProductOptionAppendSku() string {
	return productOption["appendsku"].(string)
}

// EFFECTS: Returns the price change of the product option.
func (productOption ProductOption) GetProductOptionPriceChange() string {
	return productOption["pricechange"].(string)
}

// EFFECTS: Returns the description of the product option.
func (productOption ProductOption) GetProductOptionDescription() string {
	return productOption["description"].(string)
}

// EFFECTS: Returns the invoice description of the product option.
func (productOption ProductOption) GetProductOptionInvoiceDescription() string {
	return productOption["invoice_description"].(string)
}
