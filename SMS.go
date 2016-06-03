package HologramGo

import (
	"fmt"
	"os"
)

type SMS map[string]interface{}

// REQUIRES: a device id and a phone number.
// EFFECTS: Sends an SMS to a device.
func (sms SMS) sendSMSToDevice(deviceid int, phonenumber string) {

	var params Parameters
	req := createPostRequest("/sms/incoming", params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&SMS{})

	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	//fmt.Println(SMS);
}
