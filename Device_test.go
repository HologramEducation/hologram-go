package HologramGo_Test

import (
	"testing"
	"fmt"
	"../Hologram-Go"
	"encoding/json"
	"io/ioutil"
)

func TestGetDevice(t *testing.T) {

	expectedDeviceName := "Unnamed Device (91590)"

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

	fmt.Println(device.GetDeviceName())
	deviceName := device.GetDeviceName()

	if expectedDeviceName != deviceName {
		t.Fatalf("Expected %s, got %s", expectedDeviceName, deviceName)
	}
}
