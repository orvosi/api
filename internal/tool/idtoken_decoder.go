package tool

import "google.golang.org/api/idtoken"

// IDTokenDecoder responsibles for decoding google's ID token.
type IDTokenDecoder struct {
	validator *idtoken.Validator
	audience  string
}
