package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordInserter_Executor struct {
	repo *repository.MedicalRecordInserter
	sql  sqlmock.Sqlmock
}

func TestNewMedicalRecordInserter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordInserter", func(t *testing.T) {
		exec := createMedicalRecordInserterExecutor(ctrl)
		assert.NotNil(t, exec.repo)
	})
}

func TestMedicalRecordInserter_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("can't proceed due to nil medical record", func(t *testing.T) {
		exec := createMedicalRecordInserterExecutor(ctrl)

		err := exec.repo.Insert(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyMedicalRecord, err)
	})
}

func createMedicalRecordInserterExecutor(ctrl *gomock.Controller) *MedicalRecordInserter_Executor {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Panicf("error opening a stub database connection: %v\n", err)
	}

	repo := repository.NewMedicalRecordInserter(db)
	return &MedicalRecordInserter_Executor{
		repo: repo,
		sql:  mock,
	}
}
