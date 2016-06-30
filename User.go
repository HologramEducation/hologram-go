package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

// Users is just a list of User(s).
type Users []User
type User map[string]interface{}

// REQUIRES: a device id.
// EFFECTS: Makes a HTTP Post call to create a new user.
func CreateUser(id int) User {

	var params Parameters
	req := createPostRequest("/users/", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

func GetUserAccountDetails(id int) User {

	req := createGetRequest("/users/" + strconv.Itoa(id));

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}


// REQUIRES: a new password from the user.
// EFFECTS: Changes the user's password
func ChangeUserPassword(password string) User {

	var params Parameters
	req := createPostRequest("/users/me/password/", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}


// EFFECTS: Retrieve user addresses.
func GetUserAddresses() User {

	req := createGetRequest("/users/me/addresses")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// REQUIRES: The address.
// EFFECTS: Adds a new address to the user.
func AddUserAddress() User {

	var params Parameters
	req := createPostRequest("/users/me/addresses", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// EFFECTS: Returns the user's API key.
func GetAPIKey() User {

	req := createGetRequest("/users/me/apikey")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// EFFECTS: Generates a new API key.
func GenerateNewAPIKey() User {

	var params Parameters
	req := createPostRequest("/users/me/apikey", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

///////////////////////////////////////////////////
// GENERIC USER GETTER FUNCTIONS
///////////////////////////////////////////////////

func (user User) GetUserFirstName() string {
	return user["first"].(string)
}

func (user User) GetUserLastName() string {
	return user["last"].(string)
}

func (user User) GetUserRole() string {
	return user["role"].(string)
}

// EFFECTS: Returns the user's API key.
func (user User) GetUserAPIKey() string {
	return user["apikey"].(string)
}
