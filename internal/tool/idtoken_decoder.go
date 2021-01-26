package tool

import (
	"context"

	"github.com/orvosi/api/entity"
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

// Decode decodes google token.
func (id *IDTokenDecoder) Decode(ctx context.Context, googleToken string) (*entity.User, *entity.Error) {
	payload, err := id.validator.Validate(ctx, googleToken, id.audience)
	if err != nil {
		return nil, entity.WrapError(entity.ErrUnauthorized, err.Error())
	}

	user := payloadToUser(payload)
	return user, nil
}

func payloadToUser(payload *idtoken.Payload) *entity.User {
	id, _ := payload.Claims["sub"].(string)
	name, _ := payload.Claims["name"].(string)
	email, _ := payload.Claims["email"].(string)

	return &entity.User{
		GoogleID: id,
		Name:     name,
		Email:    email,
	}
}
