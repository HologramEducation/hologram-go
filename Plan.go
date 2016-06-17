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

// EFFECTS: Returns the data plan id.
func (plan Plan) GetDataPlanId(map[string]interface{}) float64 {
	return plan["id"].(float64)
}

// EFFECTS: Returns the data plan partner id.
func (plan Plan) GetDataPlanPartnerId(map[string]interface{}) float64 {
	return plan["partnerid"].(float64)
}

// EFFECTS: Returns the data plan name.
func (plan Plan) GetDataPlanName(map[string]interface{}) string {
	return plan["name"].(string)
}

// EFFECTS: Returns the data plan description.
func (plan Plan) GetDataPlanDescription(map[string]interface{}) string {
	return plan["description"].(string)
}

// EFFECTS: Returns the data size.
func (plan Plan) GetDataPlanSize(map[string]interface{}) float64 {
	return plan["size"].(float64)
}

// EFFECTS: Returns true if it is recurring.
func (plan Plan) IsDataPlanRecurring(map[string]interface{}) bool {
	return plan["recurring"].(bool)
}

// EFFECTS: Returns true if the data plan is enabled.
func (plan Plan) IsDataPlanEnabled(map[string]interface{}) bool {
	return plan["enabled"].(bool)
}

// EFFECTS: Returns the billing period.
func (plan Plan) GetDataPlanBillingPeriod(map[string]interface{}) float64 {
	return plan["billingperiod"].(float64)
}

// EFFECTS: Returns the number of trial days left.
func (plan Plan) GetDataPlanTrialDays(map[string]interface{}) float64 {
	return plan["traildays"].(float64)
}

// EFFECTS: Returns the data plan template id.
func (plan Plan) GetDataPlanTemplateId(map[string]interface{}) float64 {
	return plan["templateid"].(float64)
}

// EFFECTS: Returns the carrier id of the data plan.
func (plan Plan) GetDataPlanCarrierId(map[string]interface{}) float64 {
	return plan["carrierid"].(float64)
}

// EFFECTS: Returns the groupid of the data plan.
func (plan Plan) GetDataPlanGroupId(map[string]interface{}) float64 {
	return plan["groupid"].(float64)
}
