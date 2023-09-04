package http

import (
	"strings"
)

// NewError creates a new Error instance.
func NewError(status Status, message ...string) IError {
	if len(message) > 0 {
		return &Error{status, strings.Join(message, " ")}
	}
	return &Error{status, status.String()}
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// StatusCode returns the HTTP status code.
func (e *Error) StatusCode() Status {
	return e.Status
}
