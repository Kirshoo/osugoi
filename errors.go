package osugoi

import "fmt"

type ErrorMessage struct {
	Type string `json:"error"`
	Description string `json:"error_description"`
	Message string `json:"message"`
}

type HttpRequestError struct {
	StatusCode int
	Message *ErrorMessage
}

func (e *HttpRequestError) Error() string {
	return fmt.Sprintf("Failed HTTP request. Status: %d. ErrorType: %s, ErrorMessage: %s", 
		e.StatusCode, e.Message.Type, e.Message.Message)
}
