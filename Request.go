package HologramGo

// This file contains all HTTP Response related behavior.

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"fmt"
)

const HOLOGRAM_REST_API_BASEURL = "https://dashboard.hologram.io/api/1/"

var err	error

// Sends a HTTP request through this instance's HTTP client.
func SendRequest(req *http.Request) (response *http.Response, err error) {

	client := &http.Client {
		//CheckRedirect: redirectPolicyFunc,
	}

	response, err = client.Do(req)

	return response, err
}

func createGetRequest(derivedUrl string) (req *http.Request) {


	req, err = http.NewRequest("GET", HOLOGRAM_REST_API_BASEURL + derivedUrl, nil)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}

// REQUIRES: the HTTP method, the endpoint of the API call, parameter names and values.
// EFFECTS: Creates and populates a HTTP request.
func createPostRequest(derivedUrl string, params Parameters) (req *http.Request) {

	data := url.Values{}
	// TODO: lolololol
	//data.Set(parameterNames, parameterValues)

	body := strings.NewReader(data.Encode())

	req, err = http.NewRequest("POST", HOLOGRAM_REST_API_BASEURL + derivedUrl, body)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
