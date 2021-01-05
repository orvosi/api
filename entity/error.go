package entity

import "fmt"

var (
	// ErrEmptyMedicalRecord indicates that a medical record is empty or null.
	ErrEmptyMedicalRecord = NewError("03-001", "MedicalRecord is empty")
)

// Error represents a data structure for error.
type Error struct {
	// Code represents error code.
	Code string `json:"code"`
	// Message represents error message.
	// This is the message that exposed to the user.
	Message string `json:"message"`
	// internalMessage represents deep error message.
	// This is should not be exposed to the user directly.
	// This attributes should be used as log.
	internalMessage string
}

// NewError creates an instance of Error.
func NewError(code, message string) *Error {
	return &Error{
		Code:            code,
		Message:         message,
		internalMessage: message,
	}
}

// Error returns internal message in one string.
func (err *Error) Error() string {
	return err.internalMessage
}

// WrapError wraps Error with given message.
// The message will be put in internalMessage attribute
// and can be accessed via Error() method.
func WrapError(err *Error, message string) *Error {
	err.internalMessage = fmt.Sprintf("%s. %s", err.internalMessage, message)
	return err
}
