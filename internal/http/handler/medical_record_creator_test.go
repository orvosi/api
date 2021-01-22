package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordCreator_Executor struct {
	handler *handler.MedicalRecordCreator
	usecase *mock_usecase.MockCreateMedicalRecord
}

func TestNewMedicalRecordCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordCreator", func(t *testing.T) {
		exec := createMedicalRecordCreatorExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func createMedicalRecordCreatorExecutor(ctrl *gomock.Controller) *MedicalRecordCreator_Executor {
	u := mock_usecase.NewMockCreateMedicalRecord(ctrl)
	h := handler.NewMedicalRecordCreator(u)
	return &MedicalRecordCreator_Executor{
		handler: h,
		usecase: u,
	}
}
