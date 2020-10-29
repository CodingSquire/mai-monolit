//Package httpserver http errors
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

// ErrorBody ...
type ErrorBody struct {
	Data      string `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type httpError struct {
	Code int
	Body ErrorBody
}

// Error returns a text message corresponding to the given error.
func (e *httpError) Error() string {
	return e.Body.ErrorText
}

// StatusCode returns an HTTP status code corresponding to the given error.
func (e *httpError) StatusCode() int {
	return e.Code
}

// ErrorProcessor ...
type ErrorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

type errorProcessor struct {
	defaultCode int
	defaultBody ErrorBody
}

//Encode writes a svc error to the given http.ResponseWriter.
func (e *errorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	code := e.defaultCode
	errorBody := e.defaultBody
	if err, ok := err.(*httpError); ok {
		if err.Code != e.defaultCode {
			code = err.Code
			errorBody = err.Body
		}
	}
	r.SetStatusCode(code)
	easyjson.MarshalToWriter(errorBody, r.BodyWriter())
	return
}

// Decode reads a Service error from the given *http.Response.
func (e *errorProcessor) Decode(r *fasthttp.Response) error {
	msgBytes := r.Body()
	var body ErrorBody
	if err := body.UnmarshalJSON(msgBytes); err != nil {
		body.Data = ""
		body.Error = true
		body.ErrorText = "error in decoding error text"
	}
	return &httpError{
		Code: r.StatusCode(),
		Body: ErrorBody{
			Data:      body.Data,
			Error:     body.Error,
			ErrorText: body.ErrorText,
		},
	}
}

// NewErrorProcessor ...
func NewErrorProcessor(defaultCode int, defaultMessage string) ErrorProcessor {
	return &errorProcessor{
		defaultCode: defaultCode,
		defaultBody: ErrorBody{
			Data:      "",
			Error:     true,
			ErrorText: defaultMessage,
		},
	}
}

// NewError ...
func NewError(status int, format string, v ...interface{}) error {
	return &httpError{
		Code: status,
		Body: ErrorBody{
			Data:      "",
			Error:     true,
			ErrorText: fmt.Sprintf(format, v...),
		},
	}
}
