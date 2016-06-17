package HologramGo

import (
	"fmt"
	"os"
)

type Plan map[string]interface{}

// Plans is just a list of Plan(s).
type Plans []Plan

// EFFECTS: Returns device data plans.
func GetDeviceDataPlans() Plan {

	req := createGetRequest("/plans/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// REQUIRES: A plan id.
// EFFECTS: Returns a given device data plan
func GetDeviceDataPlan(planid string) Plan {

	req := createGetRequest("/plans/" + string(planid))

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}
