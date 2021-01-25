package server_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	"github.com/orvosi/api/internal/http/router"
	"github.com/orvosi/api/internal/http/server"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("successfully create an instance of Server", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srv := createServer(ctrl)
		assert.NotNil(t, srv)
	})
}

func createMedicalRecordCreator(ctrl *gomock.Controller) *handler.MedicalRecordCreator {
	m := mock_usecase.NewMockCreateMedicalRecord(ctrl)
	return handler.NewMedicalRecordCreator(m)
}

func createServer(ctrl *gomock.Controller) *server.Server {
	c := createMedicalRecordCreator(ctrl)
	r := router.MedicalRecordCreator(c)
	return server.NewServer(r)
}
