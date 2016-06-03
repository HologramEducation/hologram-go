package HologramGo

// This file contains all HTTP Response related behavior.

import (
	"io"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

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

// Error returned if there was an issue parsing the response body.
type ResponseError struct {
	Body string
	Code int
}

func NewResponseError(code int, body string) ResponseError {
	return ResponseError{Code: code, Body: body}
}

func (e ResponseError) Error() string {
	return fmt.Sprintf(
		"Unable to handle response (status code %d): `%v`",
		e.Code,
		e.Body)
}


type Error map[string]interface{}

func (e Error) Code() int64 {
	return int64(e["code"].(float64))
}

func (e Error) Message() string {
	return e["message"].(string)
}

func (e Error) Error() string {
	msg := "Error %v: %v"
	return fmt.Sprintf(msg, e.Code(), e.Message())
}

type Errors map[string]interface{}

func (e Errors) Error() string {
	var (
		msg string = ""
		err Error
		ok  bool
	)
	if e["errors"] == nil {
		return msg
	}
	for _, val := range e["errors"].([]interface{}) {
		if err, ok = val.(map[string]interface{}); ok {
			msg += err.Error() + ". "
		}
	}
	return msg
}

func (e Errors) String() string {
	return e.Error()
}

func (e Errors) Errors() []Error {
	var errs = e["errors"].([]interface{})
	var out = make([]Error, len(errs))
	for i, val := range errs {
		out[i] = Error(val.(map[string]interface{}))
	}
	return out
}

type Response http.Response

func (response Response) parseBody() (b []byte, err error) {

	var (
		header string
		reader io.Reader
	)
	defer response.Body.Close()
	reader = response.Body

	b, err = ioutil.ReadAll(reader)

	return b, err
}

func (response Response) Parse(out interface{}) (err error) {

	var b []byte

	switch response.StatusCode {
	case STATUS_UNAUTHORIZED:
		fallthrough
	case STATUS_NOTFOUND:
		fallthrough
	case STATUS_GATEWAY:
		fallthrough
	case STATUS_FORBIDDEN:
		fallthrough
	case STATUS_INVALID:
		e := &Errors{}
		if b, err = response.parseBody(); err != nil {
			return
		}
		if err = json.Unmarshal(b, e); err != nil {
			err = NewResponseError(response.StatusCode, string(b))
		} else {
			err = *e
		}
		return
	case STATUS_CREATED:
		fallthrough
	case STATUS_ACCEPTED:
		fallthrough
	case STATUS_OK:
		if b, err = response.parseBody(); err != nil {
			return err
		}
		err = json.Unmarshal(b, out)
		if err == io.EOF {
			err = nil
		}
	default:
		if b, err = response.parseBody(); err != nil {
			return err
		}
		err = NewResponseError(response.StatusCode, string(b))
	}
	return err
}
