package harvest_api_client

import "testing"

func TestGetEstimate(t *testing.T) {
	a := testAPI()
	estimateResponse := mockResponse("estimates", "estimate-example.json")
	a.BaseURL = estimateResponse.URL
	estimate, err := a.GetEstimate(1439818, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if estimate == nil {
		t.Error("GetEstimate() failed.")
	}

	if estimate.ID != 1439818 {
		t.Errorf("Incorrect Estimate ID '%v'", estimate.ID)
	}
	if estimate.Subject != "Online Store - Phase 2" {
		t.Errorf("Incorrect Subject '%s'", estimate.Subject)
	}
	if *estimate.Discount != 10.0 {
		t.Errorf("Expected Discount of 10.0. Got %.1f", *estimate.Discount)
	}
	if estimate.Amount != 9630.0 {
		t.Errorf("Expected Amount of 9630.00. Got %.2f", estimate.Amount)
	}
}

func TestGetEstimates(t *testing.T) {
	a := testAPI()
	estimateResponse := mockResponse("estimates", "estimates-example.json")
	a.BaseURL = estimateResponse.URL
	estimates, err := a.GetEstimates(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if estimates == nil {
		t.Error("GetEstimates() failed.")
	}

	if len(estimates) != 2 {
		t.Errorf("Incorrect number of estimates '%v'", len(estimates))
	}
}
