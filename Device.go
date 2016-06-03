package HologramGo

import (
	"fmt"
	"os"
)

type Device map[string]interface{}

// REQUIRES: a device id.
// EFFECTS: Returns device details.
func (device Device) GetDevice(id int) {

	req := createGetRequest("/devices" + string(id))

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

func (device Device) GetDeviceDataPlan(planid string) {

	//req := createGetRequest("/devices" + string(planid))
}

// EFFECTS: Claim ownership and activate the given device.
func (device Device) ClaimOwnershipAndActiveDevice(simnumber int) {

	//var params Parameters
	//req := createPostRequest("/cellular/sim_" + string(simnumber), params)
}

func (device Device) PurchaseAndAssignPhoneNumberToDevice(deviceid int) {

	//var params Parameters
	//req := createPostRequest("/devices/" + string(deviceid) + "/addnumber", params)

}

func (device Device) GetDeviceUserId() int {
	return device["userid"].(int)
}

func (device Device) getDeviceName() string {
	return device["name"].(string)
}

// TODO: Handle as Time type
func (device Device) getWhenCreated() string {
	return device["whencreated"].(string)
}

func (device Device) getPhoneNumber() string {
	return device["phonenumber"].(string)
}
