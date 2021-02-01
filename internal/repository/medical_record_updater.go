package repository

import (
	"context"
	"database/sql"

	"github.com/orvosi/api/entity"
)

// MedicalRecordUpdater connects the database with medical record entity
// and only responsible for usecase of updating a data.
type MedicalRecordUpdater struct {
	db *sql.DB
}

// NewMedicalRecordUpdater creates an instance of MedicalRecordUpdater.
func NewMedicalRecordUpdater(db *sql.DB) *MedicalRecordUpdater {
	return &MedicalRecordUpdater{db: db}
}

// DoesRecordExist checks whether medical record which has certain id and email exists.
func (mu *MedicalRecordUpdater) DoesRecordExist(ctx context.Context, id uint64, email string) (bool, *entity.Error) {
	query := "SELECT id FROM medical_records WHERE id = $1 AND email = $2 LIMIT 1"
	row := mu.db.QueryRowContext(ctx, query, id, email)

	var tmp uint64
	err := row.Scan(&tmp)
	if err != nil && err != sql.ErrNoRows {
		return false, entity.WrapError(entity.ErrInternalServer, err.Error())
	} else if err != nil && err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
