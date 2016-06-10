package HologramGo

import (
	"fmt"
	"os"
)

type User map[string]interface{}


// REQUIRES: a device id.
// EFFECTS: Makes a HTTP Post call to create a new user.
func (user User) CreateUser(id int) {

	var params Parameters
	req := createPostRequest("/users", params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&User{})

	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(" done with User");
}

func (user User) GetUserAccountDetails(id int) {

	req := createGetRequest("/users/" + string(id));

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&User{})

	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
}


// REQUIRES: a new password from the user.
// EFFECTS: Changes the user's password
func (user User) changeUserPassword(password string) {
}


// EFFECTS: Retrieve user addresses.
func (user User) getUserAddresses() {

	//req := createGetRequest("/users/me/addresses")
}

// REQUIRES: The address.
// EFFECTS: Adds a new address to the user.
func (user User) addUserAddress() {

	//var params Parameters
	//req := createPostRequest("/users/me/addresses", params)
}

// EFFECTS: Returns the user's API key.
func (user User) getAPIKey() {

	//req := createGetRequest("/users/me/apikey")
}

// EFFECTS: Generates a new API key.
func (user User) generateNewAPIKey() {

	//var params Parameters
	//req := createPostRequest("/users/me/apikey", params)
}

func (user User) getUserFirstName() string {
	return user["first"].(string)
}

func (user User) getUserLastName() string {
	return user["last"].(string)
}

func (user User) getUserRole() string {
	return user["role"].(string)
}


