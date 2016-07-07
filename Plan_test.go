package HologramGo

import (
	"../Hologram-Go"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
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
	var plans = (payload["data"].([]interface{}))
	var plan = (HologramGo.Plan)(plans[0].(map[string]interface{}))

	// Test for data plan id
	expectedFloat := (float64)(51)
	returnedFloat := plan.GetDataPlanId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Test for data plan partner id.
	expectedFloat = (float64)(1)
	returnedFloat = plan.GetDataPlanPartnerId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Test for data plan name.
	expectedString := "1 MB"
	returnedString := plan.GetDataPlanName()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	// Test for data plan subscription
	expectedString = ""
	returnedString = plan.GetDataPlanDescription()
	if expectedString != returnedString {
		t.Fatalf("Expected %s, got %s", expectedString, returnedString)
	}

	// Test for data plan size
	expectedFloat = (float64)(1000000)
	returnedFloat = plan.GetDataPlanSize()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Test for data plan billing period
	expectedFloat = (float64)(24)
	returnedFloat = plan.GetDataPlanBillingPeriod()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}

	// Test for data plan carrier id.
	expectedFloat = (float64)(1)
	returnedFloat = plan.GetDataPlanCarrierId()
	if expectedFloat != returnedFloat {
		t.Fatalf("Expected %s, got %s", expectedFloat, returnedFloat)
	}
}
