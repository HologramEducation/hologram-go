package HologramGo

import (
	"testing"
	"fmt"
	"../Hologram-Go"
	"encoding/json"
	"io/ioutil"
)

// TODO: Add product options and categories.
func TestGetProductOptions(t *testing.T) {

	file, e := ioutil.ReadFile("json/product_options.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}

	// Test the first product
	var productOptions = (payload["data"].([]interface{}))
	var productOption = (HologramGo.ProductOption)(productOptions[1].(map[string]interface{}))

	// Check for expected product id
	expectedFloat := (float64)(2)
	returnedFloat := productOption.GetProductIdFromOption()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	expectedString := "5DASH"
	returnedString := productOption.GetProductOptionAppendSku()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "0.00"
	returnedString = productOption.GetProductOptionPriceChange()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "5 Dashes"
	returnedString = productOption.GetProductOptionInvoiceDescription()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}
}