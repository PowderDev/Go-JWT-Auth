package apiExceptions

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Cause      string `json:"cause,omitempty"`
}

func (err *HTTPError) Error() string {
	e, _ := json.Marshal(err)
	return string(e)
}

func Error500(cause string) *HTTPError {
	return &HTTPError{
		Message:    "Something went wrong on out server!",
		Cause:      cause,
		StatusCode: http.StatusInternalServerError,
	}
}

func Error400(msg string) *HTTPError {
	return &HTTPError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
	}
}

func Error401() *HTTPError {
	return &HTTPError{
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func ErrorBodyMismatch() *HTTPError {
	return &HTTPError{
		Message:    "Not all required fields were provided",
		StatusCode: http.StatusBadRequest,
	}
}
