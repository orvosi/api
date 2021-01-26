package usecase_test

import (
	"context"
	"testing"

	"github.com/orvosi/api/entity"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/orvosi/api/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordFinder_Executor struct {
	usecase  *usecase.MedicalRecordFinder
	selector *mock_usecase.MockMedicalRecordSelector
}

func TestNewMedicalRecordFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordFinder", func(t *testing.T) {
		exec := createMedicalRecordFinderExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func TestMedicalRecordFinder_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("email doesn't contain username", func(t *testing.T) {
		exec := createMedicalRecordFinderExecutor(ctrl)

		res, err := exec.usecase.FindByEmail(context.Background(), "@dummy.com")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidEmail, err)
		assert.Empty(t, res)
	})

	t.Run("email doesn't contain domain", func(t *testing.T) {
		exec := createMedicalRecordFinderExecutor(ctrl)

		res, err := exec.usecase.FindByEmail(context.Background(), "dummy@")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidEmail, err)
		assert.Empty(t, res)
	})
}

func createMedicalRecordFinderExecutor(ctrl *gomock.Controller) *MedicalRecordFinder_Executor {
	s := mock_usecase.NewMockMedicalRecordSelector(ctrl)
	u := usecase.NewMedicalRecordFinder(s)

	return &MedicalRecordFinder_Executor{
		usecase:  u,
		selector: s,
	}
}
