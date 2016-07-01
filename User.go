package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

// Users is just a list of User(s).
type Users []User
type User map[string]interface{}

// CreateUser makes a HTTP Post call to create a new user.
func CreateUser(id int) User {

	var params Parameters
	req := createPostRequest("/users/", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// GetUserAccountDetails returns the user's account details based on the given userid.
func GetUserAccountDetails(id int) User {

	req := createGetRequest("/users/" + strconv.Itoa(id))

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// ChangeUserPassword changes the user's password.
func ChangeUserPassword(password string) User {

	var params Parameters
	req := createPostRequest("/users/me/password/", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// GetUserAddresses retrieves user addresses.
func GetUserAddresses() User {

	req := createGetRequest("/users/me/addresses")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// AddUserAddress adds a new address to the user.
func AddUserAddress() User {

	var params Parameters
	req := createPostRequest("/users/me/addresses", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// GetAPIKey returns the user's API key.
func GetAPIKey() User {

	req := createGetRequest("/users/me/apikey")

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// GenerateNewAPIKey generates a new API key.
func GenerateNewAPIKey() User {

	var params Parameters
	req := createPostRequest("/users/me/apikey", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

///////////////////////////////////////////////////
// GENERIC USER GETTER FUNCTIONS
///////////////////////////////////////////////////

// GetUserFirstName returns the first name of the given user.
func (user User) GetUserFirstName() string {
	return user["first"].(string)
}

// GetUserLastName returns the last name of the given user.
func (user User) GetUserLastName() string {
	return user["last"].(string)
}

// GetUserRole returns the role of the user.
func (user User) GetUserRole() string {
	return user["role"].(string)
}

// GetUserAPIKey returns the user's API key.
func (user User) GetUserAPIKey() string {
	return user["apikey"].(string)
}
