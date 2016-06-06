package HologramGo

import (
	"fmt"
	"os"
)

type Session map[string]interface{}

// REQUIRES: a device id.
// EFFECTS: Creates a new session.
func (session Session) CreateSession(email string, password string) {

	var params Parameters
	req := createPostRequest("/auth/session", params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Session{})

	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done creating session..");
}

// TODO: func (session Session) end // without sesskey.

// REQUIRES: A session key.
// EFFECTS; Destroys a session based on the given sesskey.
func (session Session) EndSession(sesskey string) {

	var params Parameters
	req := createPostRequest("/auth/sessiondestroy", params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	err = resp.Parse(&Session{})

	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done creating session..");
}
