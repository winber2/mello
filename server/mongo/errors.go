package mongo

import "encoding/json"

const (
	melloError = "MelloError"
)

// MelloError is a standardized error struct to return when a bad request is made
type MelloError struct {
	Type    string
	Message string
}

// Error returns a MelloError which implements error
func Error(message string) error {
	return &MelloError{Type: melloError, Message: message}
}

func (m *MelloError) Error() string {
	return m.Message
}

// ToJSON is a convenience function to convert the Error into a JSON for API Responses
func (m *MelloError) ToJSON() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}
