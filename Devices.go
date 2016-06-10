package HologramGo

import (
	"fmt"
	"os"
)

// Devices is just a list of Device(s).
//type Devices []Device interface{}

type Devices map[string]interface{}

// REQUIRES: a device id.
// EFFECTS: Returns device details.
func (devices Devices) GetDevices() {

	req := createGetRequest("/devices")

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Devices{})
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}
