package HologramGo

// This file contains all HTTP Response related behavior.

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTP status codes
const (
	STATUS_OK           = 200
	STATUS_CREATED      = 201
	STATUS_ACCEPTED     = 202
	STATUS_NO_CONTENT   = 204
	STATUS_INVALID      = 400
	STATUS_UNAUTHORIZED = 401
	STATUS_FORBIDDEN    = 403
	STATUS_NOTFOUND     = 404
	STATUS_LIMIT        = 429
	STATUS_GATEWAY      = 502
)

// EFFECTS: Takes in the response and unmarshalls it into the appropriate interface.
func unmarshallIntoObject(resp *Response) map[string]interface{} {

	var payload = Placeholder{}
	err = resp.parsePayload(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// EFFECTS: Takes in the response and unmarshalls into an array object.
func unmarshallIntoArrayObject(resp *Response) []interface{} {

	var payload = Placeholder{}
	err = resp.parsePayload(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].([]interface{})
}

type Response http.Response

func (response Response) parseBody() (b []byte, err error) {

	var (
		//header string
		reader io.Reader
	)

	defer response.Body.Close()
	reader = response.Body

	b, err = ioutil.ReadAll(reader)

	return b, err
}

// EFFECTS: Populates the given object based on the returned response
func (response Response) parsePayload(out interface{}) (err error) {

	var b []byte

	statusCode := response.StatusCode

	if statusCode == STATUS_UNAUTHORIZED || statusCode == STATUS_NOTFOUND ||
		statusCode == STATUS_GATEWAY || statusCode == STATUS_FORBIDDEN ||
		statusCode == STATUS_INVALID {

		e := &Errors{}
		b, err = response.parseBody()
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, e)

		if err != nil {
			err = NewResponseError(response.StatusCode, string(b))
		} else {
			err = *e
		}
	} else if statusCode == STATUS_CREATED || statusCode == STATUS_ACCEPTED ||
		statusCode == STATUS_OK {
		b, err = response.parseBody()
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, out)
		if err == io.EOF {
			err = nil
		}
	} else {
		b, err = response.parseBody()
		if err != nil {
			return err
		}
		err = NewResponseError(response.StatusCode, string(b))
	}
	return err
}
