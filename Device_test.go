package HologramGo

import "testing"

func TestGetDevice(t *testing.T) {
	expectedDeviceName := "SIM Card"

	var device = GetDevice(1)

	deviceName := device.GetDeviceName()

	if expectedDeviceName != deviceName {
		t.Fatalf("Expected %s, got %s", expectedDeviceName, deviceName)
	}
}
