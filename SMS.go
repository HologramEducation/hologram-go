package HologramGo

import (
	"fmt"
	"os"
)

// SMS implements the SMS type returned from the response.
type SMS map[string]interface{}

// SendSMSToDevice sends an SMS to a device and returns the response.
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
