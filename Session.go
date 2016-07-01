package HologramGo

import (
	"fmt"
	"os"
)

// Session implements the Session type returned from the response.
type Session map[string]interface{}

// CreateSession creates a new session.
func CreateSession(email string, password string) Session {

	var params Parameters
	req := createPostRequest("/auth/session", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}

// TODO: func (session Session) end // without sesskey.

// EndSession destroys a session based on the given sesskey.
func EndSession(sesskey string) Session {

	var params Parameters
	req := createPostRequest("/auth/sessiondestroy", params)

	resp, err := sendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	return unmarshallIntoObject(resp)
}
