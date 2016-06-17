package HologramGo

import (
	"testing"
	"fmt"
	"../Hologram-Go"
	"encoding/json"
	"io/ioutil"
)

func TestGetDeviceDataPlan(t *testing.T) {

	file, e := ioutil.ReadFile("json/plan.json")
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	var payload = HologramGo.Placeholder{}
	err := json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Unable to parse file")
	}
	var plan = (HologramGo.Plan)(payload["data"].(map[string]interface{}))

	// Test for data plan id
	expectedFloat := (float64)(51)
	returnedFloat := plan.GetDataPlanId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}
}
