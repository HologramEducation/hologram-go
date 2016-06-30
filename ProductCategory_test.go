package HologramGo

import (
	"testing"
	"fmt"
	"../Hologram-Go"
	"encoding/json"
	"io/ioutil"
)

func TestGetProductCategories(t *testing.T) {

	file, e := ioutil.ReadFile("json/product_categories.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}

	// Test the first product
	var productCategories = (payload["data"].([]interface{}))
	var productCategory = (HologramGo.ProductCategory)(productCategories[0].(map[string]interface{}))

	// Check for expected category.
	expectedString := "accessories"
	returnedString := productCategory.GetProductCategoryName()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	// Check for expected category.
	//expectedIntSlice := []int{15, 19, 11, 10, 9, 8, 7, 23}
	//returnedIntSlice := productCategory.GetProductIdsFromCategory()

	//// Length check
	//if len(expectedIntSlice) != len(returnedIntSlice) {
	//	t.Fatalf("Expected length %s, got %s", expectedString, returnedString)
	//}

	//for i := range expectedIntSlice {

	//	if expectedIntSlice[i] != returnedIntSlice[i] {
	//		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	//	}
	//}
}
