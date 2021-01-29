package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordFinder_Executor struct {
	handler *handler.MedicalRecordFinder
	usecase *mock_usecase.MockFindMedicalRecord
}

func TestNewMedicalRecordFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordFinder", func(t *testing.T) {
		exec := createMedicalRecordFinderExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func createMedicalRecordFinderExecutor(ctrl *gomock.Controller) *MedicalRecordFinder_Executor {
	u := mock_usecase.NewMockFindMedicalRecord(ctrl)
	h := handler.NewMedicalRecordFinder(u)
	return &MedicalRecordFinder_Executor{
		handler: h,
		usecase: u,
	}
}
