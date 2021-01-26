package tool

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/api/idtoken"
)

// IDTokenDecoder responsibles for decoding google's ID token.
type IDTokenDecoder struct {
	validator *idtoken.Validator
	audience  string
}

// NewIDTokenDecoder creates an instance of IDTokenDecoder.
func NewIDTokenDecoder(audience string, credential []byte) (*IDTokenDecoder, error) {
	v, err := idtoken.NewValidator(
		context.Background(),
		idtoken.WithCredentialsJSON(credential),
	)
	if err != nil {
		return nil, errors.Wrap(err, "[NewIDTokenDecoder] fail to create IDTokenDecoder")
	}

	return &IDTokenDecoder{
		validator: v,
		audience:  audience,
	}, nil
}
