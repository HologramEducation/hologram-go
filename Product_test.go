package HologramGo

import (
	"github.com/hologram-io/hologram-go"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// func TestGetProducts(t *testing.T) {

// 	file, e := ioutil.ReadFile("json/device.json")
// 	if e != nil {
// 		fmt.Printf("Error opening file: %v\n", e)
// 	}

// 	var payload = HologramGo.Placeholder{}
// 	err := json.Unmarshal(file, &payload)
// 	if err != nil {
// 		fmt.Println("Unable to parse file")
// 	}
// 	var device = (HologramGo.Device)(payload["data"].(map[string]interface{}))

// 	// Check for expected device id
// 	expectedFloat := (float64)(1)
// 	returnedFloat := device.GetDeviceId()
// 	if expectedFloat != returnedFloat {
// 		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
// 	}

// }

func TestGetProduct(t *testing.T) {

	file, e := ioutil.ReadFile("json/product.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}
	var product = (HologramGo.Product)(payload["data"].(map[string]interface{}))

	// Check for expected product id
	expectedFloat := (float64)(1)
	returnedFloat := product.GetProductId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	expectedString := "SIM"
	returnedString := product.GetProductSku()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "SIM Card"
	returnedString = product.GetProductName()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "The Hologram Global SIM Card. Comes in a variety of sizes for any application. (Pick your data plan when you go to activate the SIM.)"
	returnedString = product.GetProductDescription()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "5.00"
	returnedString = product.GetProductPrice()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "https://dashboard.hologram.io/img/sim_card.jpg"
	returnedString = product.GetProductImageUrl()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	expectedString = "Global SIM Card"
	returnedString = product.GetProductInvoiceDescription()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}
}
