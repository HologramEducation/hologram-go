package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

// Device object returned in the response.
type Device map[string]interface{}

// Devices is just a list of Device(s).
type Devices []interface{}

// GetDevices returns device details.
func GetDevices() Devices {

	req := createGetRequest("/devices")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

// GetDevice returns a device detail based on the given deviceid.
func GetDevice(deviceid int) Device {

	req := createGetRequest("/devices/" + strconv.Itoa(deviceid))

	resp, err := sendRequest(req)

	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// ClaimOwnershipAndActivateDevice claims ownership and activate the given device.
func ClaimOwnershipAndActivateDevice(simnumber int) Device {

	var params Parameters
	req := createPostRequest("/cellular/sim_"+strconv.Itoa(simnumber)+"/claim", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// PurchaseAndAssignPhoneNumberToDevice purchases and assigns a phone number to the device.
func PurchaseAndAssignPhoneNumberToDevice(deviceid int) Device {

	var params Parameters
	req := createPostRequest("/devices/"+strconv.Itoa(deviceid)+"/addnumber", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

///////////////////////////////////////////////////
// GENERIC DEVICE GETTER FUNCTIONS
///////////////////////////////////////////////////

// GetDeviceId returns the id.
func (device Device) GetDeviceId() float64 {
	return device["id"].(float64)
}

// GetDeviceUserId returns the user id.
func (device Device) GetDeviceUserId() float64 {
	return device["userid"].(float64)
}

// GetDeviceName returns the device name.
func (device Device) GetDeviceName() string {
	return device["name"].(string)
}

// GetDeviceType returns the device type.
func (device Device) GetDeviceType() string {
	return device["type"].(string)
}

// GetWhenCreated returns a UNIX timestamp of the creation time.
func (device Device) GetWhenCreated() string {
	return device["whencreated"].(string)
}

// GetPhoneNumber returns a phone number.
func (device Device) GetPhoneNumber() string {
	return device["phonenumber"].(string)
}

// GetTunnelable returns true if it is tunnelable.
func (device Device) GetTunnelable() bool {
	return device["tunnelable"].(bool)
}

// TODO: links/cellular
