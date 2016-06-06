package HologramGo

import (
	"fmt"
	"os"
)

type Device map[string]interface{}

// REQUIRES: a device id.
// EFFECTS: Returns device details.
func (device Device) GetDevice(deviceid int) {

	req := createGetRequest("/devices/" + string(deviceid))

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

	fmt.Println(" done with Device");
}

// REQUIRES: A plan id.
// EFFECTS: Returns a given device data plan
func (device Device) GetDeviceDataPlan(planid string) {

	req := createGetRequest("/devices/" + string(planid))

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

	fmt.Println(" done with Device");
}

// REQUIRES: A sim number.
// EFFECTS: Claim ownership and activate the given device.
func (device Device) ClaimOwnershipAndActiveDevice(simnumber int) {

	var params Parameters
	req := createPostRequest("/cellular/sim_" + string(simnumber), params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Device{})
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(" done with Device");
}

// REQUIRES: A device id.
// EFFECTS: Purchase and assign a phone number to the device.
func (device Device) PurchaseAndAssignPhoneNumberToDevice(deviceid int) {

	//var params Parameters
	//req := createPostRequest("/devices/" + string(deviceid) + "/addnumber", params)

}

// EFFECTS: Returns the user id.
func (device Device) GetDeviceUserId() int {
	return device["userid"].(int)
}

// EFFECTS: Returns the device name.
func (device Device) GetDeviceName() string {
	return device["name"].(string)
}

// TODO: Handle as Time type
func (device Device) GetWhenCreated() string {
	return device["whencreated"].(string)
}

// EFFECTS: Returns a phone number.
func (device Device) GetPhoneNumber() string {
	return device["phonenumber"].(string)
}
