package HologramGo

// This file contains all Error types used within the package.

import (
	"fmt"
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

