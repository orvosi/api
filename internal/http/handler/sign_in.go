package handler

import "github.com/orvosi/api/usecase"

// Signer handles HTTP request and response
// for sign in.
type Signer struct {
	signin usecase.SignIn
}

// NewSigner creates an instance of Medical.
func NewSigner(signin usecase.SignIn) *Signer {
	return &Signer{
		signin: signin,
	}
}
