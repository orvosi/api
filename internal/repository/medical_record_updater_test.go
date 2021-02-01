package repository_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/orvosi/api/internal/repository"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordUpdater_Executor struct {
	repo *repository.MedicalRecordUpdater
	sql  sqlmock.Sqlmock
}

func TestNewMedicalRecordUpdater(t *testing.T) {
	t.Run("successfully create an instance of MedicalRecordUpdater", func(t *testing.T) {
		exec := createMedicalRecordUpdaterExecutor()
		assert.NotNil(t, exec.repo)
	})
}

func createMedicalRecordUpdaterExecutor() *MedicalRecordUpdater_Executor {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Panicf("error opening a stub database connection: %v\n", err)
	}

	repo := repository.NewMedicalRecordUpdater(db)
	return &MedicalRecordUpdater_Executor{
		repo: repo,
		sql:  mock,
	}
}
