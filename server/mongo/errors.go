package mongo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	melloError = "MelloError"
	formError  = "FormError"
)

// Error is a standardized error struct to return when a bad request is made
type Error struct {
	Type       string
	Message    string
	Resolution string
}

type FormError struct {
	Type     string
	Message  string
	Failures []string
}

// CreateError returns a Error which implements error
func CreateError(message string) error {
	return &Error{Type: melloError, Message: message}
}

func (m *Error) Error() string {
	return m.Message
}

// ToJSON is a convenience function to convert the Error into a JSON for API Responses
func (m *Error) ToJSON() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// CreateFormError creates a FormError used for bad form requests
func CreateFormError(model string, keys []string) error {
	var errors []string
	var message string

	for _, key := range keys {
		errors = append(errors, getMessage(model, key))
		message = strings.Join(errors, ", ")
	}

	return &FormError{
		Type:     formError,
		Message:  message,
		Failures: keys,
	}
}

func (f *FormError) Error() string {
	return f.Message
}

func getMessage(model string, key string) string {
	return fmt.Sprintf("%v with that %v already exists", model, key)
}
