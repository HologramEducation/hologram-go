
package HologramGo

import (
	"fmt"
	"os"
)

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
