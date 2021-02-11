package entity

import "fmt"

var (
	// ErrInternalServer indicates there is unexpected problem occurs in the system itself.
	// The detail of the error/problem should be known in internal message.
	ErrInternalServer = NewError("01-001", "Internal server error")
	// ErrUnauthorized is returned when a request doesn't include authorization in its header.
	// The authorization must be using bearer authorization.
	// It also can be returned if the authorization is invalid.
	ErrUnauthorized = NewError("01-002", "Request is unauthorized")
	// ErrInvalidGoogleToken is returned when the id token is invalid,
	// whether it has expired or it is not google id token.
	ErrInvalidGoogleToken = NewError("01-003", "Google ID Token is invalid")
	// ErrInvalidID is returned when the entity id can't be decoded or invalid.
	ErrInvalidID = NewError("01-004", "Entity ID is invalid")
	// ErrWrongContentType is returned when content-type in request's header is not as expected.
	ErrWrongContentType = NewError("01-005", "Wrong content type")

	// ErrEmptyMedicalRecord indicates that a medical record is empty or null.
	ErrEmptyMedicalRecord = NewError("02-001", "MedicalRecord is empty")
	// ErrInvalidMedicalRecordAttribute indicates that a medical record is empty or null.
	ErrInvalidMedicalRecordAttribute = NewError("02-002", "Medical record's attributes are invalid. Please, check all attributes")
	// ErrInvalidMedicalRecordRequest indicates that a medical record request that is sent over HTTP is invalid.
	ErrInvalidMedicalRecordRequest = NewError("02-003", "Medical record request is invalid. Please, check the JSON request")
	// ErrInvalidEmail indicates that the email passed is invalid.
	// The validation is run using regex "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	// and LookupMX from net package.
	ErrInvalidEmail = NewError("02-004", "Email is invalid. Please, check the email")
	// ErrMedicalRecordNotFound indicates that the medical record can't be found.
	ErrMedicalRecordNotFound = NewError("02-005", "Medical record not found")

	// ErrEmptyUser indicates that a user is empty or null.
	ErrEmptyUser = NewError("03-001", "User is empty")
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
