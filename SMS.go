package HologramGo

import (
	"fmt"
	"os"
)

type SMS map[string]interface{}

// REQUIRES: a device id and a phone number.
// EFFECTS: Sends an SMS to a device and returns the response
func SendSMSToDevice(deviceid int, phonenumber string) SMS {

	var params Parameters
	req := createPostRequest("/sms/incoming", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}
