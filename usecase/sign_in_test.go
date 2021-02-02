package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/orvosi/api/usecase"
	"github.com/stretchr/testify/assert"
)

type SignIn_Executor struct {
	usecase *usecase.Signer
	repo    *mock_usecase.MockInsertUserRepository
}

func TestNewSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of SignIn", func(t *testing.T) {
		exec := createSignInExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func createSignInExecutor(ctrl *gomock.Controller) *SignIn_Executor {
	r := mock_usecase.NewMockInsertUserRepository(ctrl)
	u := usecase.NewSigner(r)

	return &SignIn_Executor{
		usecase: u,
		repo:    r,
	}
}
