package HologramGo

// This file contains all HTTP Response related behavior.

import (
	"golang.org/x/net/publicsuffix"
	"net/http/cookiejar"
	"net/http"
	"net/url"
	"os"
	"strings"
	"fmt"
	"log"
)

const HOLOGRAM_REST_API_BASEURL = "https://dashboard.hologram.io/api/1"

var err	error

// These two variables will store user's username and password.
var username string
var password string

// REQUIRES: User's username and password.
// EFFECTS: Initializes the user's username and password state to be used in each
//			basic authenticated API call.
func InitializeUsernameAndPassword(input_username string, input_password string) {
	username, password = input_username, input_password
}

// EFFECTS: Sends a HTTP request through this instance's HTTP client.
func SendRequest(req *http.Request) (response *Response, err error) {

	options := cookiejar.Options {
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client {
		//CheckRedirect: redirectPolicyFunc,
		Jar: jar,
	}

	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)

	response = (*Response)(resp)
	return response, err
}

// REQUIRES: The endpoint of the API call.
// EFFECTS: Creates and populates a HTTP GET request.
func createGetRequest(derivedUrl string) (req *http.Request) {

	req, err = http.NewRequest("GET", HOLOGRAM_REST_API_BASEURL + derivedUrl, nil)
	fmt.Println(HOLOGRAM_REST_API_BASEURL + derivedUrl);
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}

// REQUIRES: The endpoint of the API call, parameter names and values.
// EFFECTS: Creates and populates a HTTP POST request.
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
