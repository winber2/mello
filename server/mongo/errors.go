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

func toJSON(j interface{}) (string, error) {
	b, err := json.Marshal(j)
	return string(b), err
}

// MelloError is a standardized error for any server side issues
type MelloError struct {
	Type       string
	Message    string
	Resolution string
}

// CreateError returns a Error for any server side issues
func CreateError(message string) error {
	return &MelloError{Type: melloError, Message: message}
}

// Error converts the the error into a string -- implementing the error interface
func (m *MelloError) Error() string {
	return m.Message
}

// ToJSON is a convenience function to convert the Error into a JSON for API Responses
func (m *MelloError) ToJSON() (string, error) {
	return toJSON(m)
}

// FormError is a standardized error for bad form requests
type FormError struct {
	Type     string
	Message  string
	Failures []string
}

func getMessage(model string, key string) string {
	return fmt.Sprintf("%v with that %v already exists", model, key)
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

// Error converts the the error into a string -- implementing the error interface
func (f *FormError) Error() string {
	return f.Message
}

// ToJSON is a convenience function to convert the Error into a JSON for API Responses
func (f *FormError) ToJSON() (string, error) {
	return toJSON(f)
}
