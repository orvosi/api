package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type Signer_Executor struct {
	handler *handler.Signer
	usecase *mock_usecase.MockSignIn
}

func TestNewSigner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of Signer", func(t *testing.T) {
		exec := createSignerExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func createSignerExecutor(ctrl *gomock.Controller) *Signer_Executor {
	u := mock_usecase.NewMockSignIn(ctrl)
	h := handler.NewSigner(u)
	return &Signer_Executor{
		handler: h,
		usecase: u,
	}
}
