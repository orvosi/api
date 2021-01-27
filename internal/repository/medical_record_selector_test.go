package repository_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/orvosi/api/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordSelector_Executor struct {
	repo *repository.MedicalRecordSelector
	sql  sqlmock.Sqlmock
}

func TestNewMedicalRecordSelector(t *testing.T) {
	t.Run("successfully create an instance of MedicalRecordSelector", func(t *testing.T) {
		exec := createMedicalRecordSelectorExecutor()
		assert.NotNil(t, exec.repo)
	})
}

func createMedicalRecordSelectorExecutor() *MedicalRecordSelector_Executor {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Panicf("[createMedicalRecordSelectorExecutor] error opening a stub database connection: %v\n", err)
	}

	repo := repository.NewMedicalRecordSelector(db)
	return &MedicalRecordSelector_Executor{
		repo: repo,
		sql:  mock,
	}
}
