package repository_test

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/hashids"
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

	t.Run("query doesn't return inserted id", func(t *testing.T) {
		exec := createMedicalRecordInserterExecutor(ctrl)
		record := createValidMedicalRecord()

		exec.sql.ExpectQuery(`INSERT INTO medical_records \(user_id, symptom, diagnosis, therapy, created_at, updated_at, created_by, updated_by\)`).
			WillReturnError(errors.New("fail to insert to database"))

		err := exec.repo.Insert(context.Background(), record)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
	})

	t.Run("successfully insert a new medical record", func(t *testing.T) {
		exec := createMedicalRecordInserterExecutor(ctrl)
		record := createValidMedicalRecord()

		exec.sql.ExpectQuery(`INSERT INTO medical_records \(user_id, symptom, diagnosis, therapy, created_at, updated_at, created_by, updated_by\)`).
			WillReturnRows(sqlmock.
				NewRows([]string{"id"}).
				AddRow(999),
			)

		err := exec.repo.Insert(context.Background(), record)

		assert.Nil(t, err)
		assert.Equal(t, hashids.ID(999), record.ID)
	})
}

func createValidMedicalRecord() *entity.MedicalRecord {
	return &entity.MedicalRecord{
		ID:        hashids.ID(1),
		Symptom:   "symptom",
		Diagnosis: "diagnosis",
		Therapy:   "therapy",
		Result:    "result",
		User: &entity.User{
			ID:       hashids.ID(1),
			Email:    "email@provider.com",
			Name:     "User 1",
			GoogleID: "super-long-google-id",
		},
	}
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
