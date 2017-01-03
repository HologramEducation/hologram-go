package HologramGo

import "testing"

func TestGetProduct(t *testing.T) {
	expectedProductName := "SIM Card"

	var product = GetProduct(1)

	productName := product.GetProductName()

	if expectedProductName != productName {
		t.Fatalf("Expected %s, got %s", expectedProductName, product.GetProductName())
	}
}
