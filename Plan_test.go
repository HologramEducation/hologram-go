package HologramGo

import "testing"

func TestGetDeviceDataPlan(t *testing.T) {
	expectedDeviceName := "SIM Card"

	var plan = GetDeviceDataPlan(1)

	deviceName := plan.GetDeviceName()

	if expectedDeviceName != deviceName {
		t.Fatalf("Expected %s, got %s", expectedDeviceName, deviceName)
	}
}
