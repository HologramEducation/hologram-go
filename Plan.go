package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

// Plan is basically the returned Plan type in the response.
type Plan map[string]interface{}

// Plans is just a list of Plan(s).
type Plans []interface{}

// GetDeviceDataPlans returns device data plans.
func GetDeviceDataPlans() Plans {

	req := createGetRequest("/plans/")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoArrayObject(resp)
}

// GetDeviceDataPlan returns a given device data plan.
func GetDeviceDataPlan(planid int) Plan {

	req := createGetRequest("/plans/" + strconv.Itoa(planid))

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	plans := unmarshallIntoArrayObject(resp)

	return (Plan)(plans[0].(map[string]interface{}))
}

///////////////////////////////////////////////////
// GENERIC PLAN GETTER FUNCTIONS
///////////////////////////////////////////////////

// GetDataPlanId returns the data plan id.
func (plan Plan) GetDataPlanId() float64 {
	return plan["id"].(float64)
}

// GetDataPlanPartnerId returns the data plan partner id.
func (plan Plan) GetDataPlanPartnerId() float64 {
	return plan["partnerid"].(float64)
}

// GetDataPlanName returns the data plan name.
func (plan Plan) GetDataPlanName() string {
	return plan["name"].(string)
}

// GetDataPlanDescription returns the data plan description.
func (plan Plan) GetDataPlanDescription() string {
	return plan["description"].(string)
}

// GetDataPlanSize returns the data size.
func (plan Plan) GetDataPlanSize() float64 {
	return plan["data"].(float64)
}

// IsDataPlanRecurring returns true if it is recurring.
func (plan Plan) IsDataPlanRecurring() bool {
	return plan["recurring"].(bool)
}

// IsDataPlanEnabled returns true if the data plan is enabled.
func (plan Plan) IsDataPlanEnabled() bool {
	return plan["enabled"].(bool)
}

// GetDataPlanBillingPeriod returns the billing period.
func (plan Plan) GetDataPlanBillingPeriod() float64 {
	return plan["billingperiod"].(float64)
}

// GetDataPlanTrialDays returns the number of trial days left.
func (plan Plan) GetDataPlanTrialDays() float64 {
	return plan["trialdays"].(float64)
}

// GetDataPlanTemplateId returns the data plan template id.
func (plan Plan) GetDataPlanTemplateId() float64 {
	return plan["templateid"].(float64)
}

// GetDataPlanCarrierId returns the carrier id of the data plan.
func (plan Plan) GetDataPlanCarrierId() float64 {
	return plan["carrierid"].(float64)
}

// GetDataPlanGroupId returns the groupid of the data plan.
func (plan Plan) GetDataPlanGroupId() float64 {
	return plan["groupid"].(float64)
}
