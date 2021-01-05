package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/orvosi/api/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordCreator_Executor struct {
	usecase  *usecase.MedicalRecordCreator
	inserter *mock_usecase.MockMedicalRecordInserter
}

func TestNewMedicalRecordCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordCreator", func(t *testing.T) {
		exec := createMedicalRecordCreatorExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func createMedicalRecordCreatorExecutor(ctrl *gomock.Controller) *MedicalRecordCreator_Executor {
	i := mock_usecase.NewMockMedicalRecordInserter(ctrl)
	u := usecase.NewMedicalRecordCreator(i)

	return &MedicalRecordCreator_Executor{
		usecase:  u,
		inserter: i,
	}
}
