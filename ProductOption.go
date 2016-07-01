package HologramGo

import (
	"fmt"
	"os"
)

// ProductOption implements the product option returned from the response.
type ProductOption map[string]interface{}

// ProductOptions is just a list of ProductOption(s)
type ProductOptions []interface{}

// GetProductOptions returns product options.
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

// GetProductIdFromOption returns the product id.
func (productOption ProductOption) GetProductIdFromOption() float64 {
	return productOption["productid"].(float64)
}

// GetProductOptionAppendSku returns the sku of the product option.
func (productOption ProductOption) GetProductOptionAppendSku() string {
	return productOption["appendsku"].(string)
}

// GetProductOptionPriceChang returns the price change of the product option.
func (productOption ProductOption) GetProductOptionPriceChange() string {
	return productOption["pricechange"].(string)
}

// GetProductOptionDescription returns the description of the product option.
func (productOption ProductOption) GetProductOptionDescription() string {
	return productOption["description"].(string)
}

// GetProductOptionInvoiceDescription returns the invoice description of the product option.
func (productOption ProductOption) GetProductOptionInvoiceDescription() string {
	return productOption["invoice_description"].(string)
}
