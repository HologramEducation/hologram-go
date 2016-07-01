package HologramGo_Test

import (
	"testing"
	"fmt"
	"../Hologram-Go"
	"encoding/json"
	"io/ioutil"
)

func TestGetDevices(t *testing.T) {

	file, e := ioutil.ReadFile("json/devices.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}

	var devices = payload["data"].([]interface{})
	var device = (HologramGo.Device)(devices[0].(map[string]interface{}))

	// Check for expected device id
	expectedFloat := (float64)(84)
	returnedFloat := device.GetDeviceId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Check for expected device user id
	expectedFloat = (float64)(4477)
	returnedFloat = device.GetDeviceUserId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Check for expected device name
	expected := "Unnamed Device (90123)"
	returned := device.GetDeviceName()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device type
	expected = "Unknown"
	returned = device.GetDeviceType()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device creation date
	expected = "2016-06-06 21:51:16"
	returned = device.GetWhenCreated()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device phone number
	expected = ""
	returned = device.GetPhoneNumber()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}
}

func TestGetDevice(t *testing.T) {

	file, e := ioutil.ReadFile("json/device.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}
	var device = (HologramGo.Device)(payload["data"].(map[string]interface{}))

	// Check for expected device id
	expectedFloat := (float64)(1)
	returnedFloat := device.GetDeviceId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Check for expected device user id
	expectedFloat = (float64)(231)
	returnedFloat = device.GetDeviceUserId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Check for expected device name
	expected := "Unnamed Device (91590)"
	returned := device.GetDeviceName()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device type
	expected = "Unknown"
	returned = device.GetDeviceType()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device creation date
	expected = "2015-03-25 03:12:13"
	returned = device.GetWhenCreated()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}

	// Check for expected device phone number
	expected = ""
	returned = device.GetPhoneNumber()
	if expected != returned {
		t.Fatalf("Expected %s, got %s", expected, returned)
	}
}
