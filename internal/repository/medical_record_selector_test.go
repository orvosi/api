package repository_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/orvosi/api/entity"
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

func TestMedicalRecordSelector_FindByEmail(t *testing.T) {
	t.Run("select query returns error", func(t *testing.T) {
		exec := createMedicalRecordSelectorExecutor()

		exec.sql.ExpectQuery(`SELECT id, symptom, diagnosis, therapy, result, updated_at FROM medical_records WHERE email = \$1 ORDER BY id ASC`).
			WillReturnError(errors.New("fail to select from database"))

		res, err := exec.repo.FindByEmail(context.Background(), "dummy@dummy.com")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Empty(t, res)
	})

	t.Run("row scan returns error", func(t *testing.T) {
		exec := createMedicalRecordSelectorExecutor()

		exec.sql.ExpectQuery(`SELECT id, symptom, diagnosis, therapy, result, updated_at FROM medical_records WHERE email = \$1 ORDER BY id ASC`).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "symptom", "diagnosis", "therapy", "result", "updated_at"}).
				AddRow(1, "Symptom", "Diagnosis", "Therapy", "Result", time.Now()).
				AddRow(2, "Symptom", "Diagnosis", "Therapy", "Result", "time.Now()"),
			)

		res, err := exec.repo.FindByEmail(context.Background(), "dummy@dummy.com")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, 1, len(res))
	})

	t.Run("successfully retrieve all rows", func(t *testing.T) {
		exec := createMedicalRecordSelectorExecutor()

		exec.sql.ExpectQuery(`SELECT id, symptom, diagnosis, therapy, result, updated_at FROM medical_records WHERE email = \$1 ORDER BY id ASC`).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "symptom", "diagnosis", "therapy", "result", "updated_at"}).
				AddRow(1, "Symptom", "Diagnosis", "Therapy", "Result", time.Now()).
				AddRow(2, "Symptom", "Diagnosis", "Therapy", "Result", time.Now()),
			)

		res, err := exec.repo.FindByEmail(context.Background(), "dummy@dummy.com")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, 2, len(res))
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
