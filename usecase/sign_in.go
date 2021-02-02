package usecase

import (
	"context"

	"github.com/orvosi/api/entity"
)

// SignIn defines the business logic to sign in.
type SignIn interface {
	// SignIn signs a user in to the system.
	SignIn(ctx context.Context, user *entity.User) error
}

// Signer responsibles for sign-in workflow.
type Signer struct{}
