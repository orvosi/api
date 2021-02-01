package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/orvosi/api/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordUpdater_Executor struct {
	usecase *usecase.MedicalRecordUpdater
	repo    *mock_usecase.MockUpdateMedicalRecordRepository
}

func TestNewMedicalRecordUpdater(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordUpdater", func(t *testing.T) {
		exec := createMedicalRecordUpdaterExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func createMedicalRecordUpdaterExecutor(ctrl *gomock.Controller) *MedicalRecordUpdater_Executor {
	r := mock_usecase.NewMockUpdateMedicalRecordRepository(ctrl)
	u := usecase.NewMedicalRecordUpdater(r)

	return &MedicalRecordUpdater_Executor{
		usecase: u,
		repo:    r,
	}
}
