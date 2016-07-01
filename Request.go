package HologramGo

// This file contains all HTTP Response related behavior.

import (
	"golang.org/x/net/publicsuffix"
	"net/http/cookiejar"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"
	"encoding/json"
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
func InitializeUsernameAndPassword(credentialFile string) {

	file, e := ioutil.ReadFile(credentialFile)
	if e != nil {
		fmt.Printf("Error opening file: %v\n", e)
	}

	type Payload map[string]interface{}
	var jsonpayload = Payload{}
	json.Unmarshal(file, &jsonpayload)

	username = jsonpayload["username"].(string)
	password = jsonpayload["password"].(string)
}

// EFFECTS: Sends a HTTP request through this instance's HTTP client.
func sendRequest(req *http.Request) (response *Response, err error) {

	options := cookiejar.Options {
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client {
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

	// Set all body parameters here.
	for i, _ := range params.items {
		data.Set(params.items[i], params.values[i])
	}

	body := strings.NewReader(data.Encode())

	req, err = http.NewRequest("POST", HOLOGRAM_REST_API_BASEURL + derivedUrl, body)
	fmt.Println(HOLOGRAM_REST_API_BASEURL + derivedUrl)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
