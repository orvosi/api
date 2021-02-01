package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordUpdater_Executor struct {
	handler *handler.MedicalRecordUpdater
	usecase *mock_usecase.MockUpdateMedicalRecord
}

func TestNewMedicalRecordUpdater(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordUpdater", func(t *testing.T) {
		exec := createMedicalRecordUpdaterExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func createMedicalRecordUpdaterExecutor(ctrl *gomock.Controller) *MedicalRecordUpdater_Executor {
	u := mock_usecase.NewMockUpdateMedicalRecord(ctrl)
	h := handler.NewMedicalRecordUpdater(u)
	return &MedicalRecordUpdater_Executor{
		handler: h,
		usecase: u,
	}
}
