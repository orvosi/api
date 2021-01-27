package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/orvosi/api/entity"
)

// MedicalRecordSelector connects the database with medical record entity
// and only responsible for retrieving medical record data.
type MedicalRecordSelector struct {
	db *sql.DB
}

// NewMedicalRecordSelector creates an instance of MedicalRecordSelector.
func NewMedicalRecordSelector(db *sql.DB) *MedicalRecordSelector {
	return &MedicalRecordSelector{db: db}
}

// FindByEmail finds all medical records bounded to specific email.
func (ms *MedicalRecordSelector) FindByEmail(ctx context.Context, email string) ([]*entity.MedicalRecord, *entity.Error) {
	query := "SELECT id, symptom, diagnosis, therapy, result, updated_at FROM medical_records WHERE email = $1 ORDER BY id ASC"
	rows, err := ms.db.QueryContext(ctx, query, email)
	if err != nil {
		return []*entity.MedicalRecord{}, entity.WrapError(entity.ErrInternalServer, err.Error())
	}
	defer rows.Close()

	var result []*entity.MedicalRecord
	for rows.Next() {
		var tmp entity.MedicalRecord
		if err := rows.Scan(&tmp.ID, &tmp.Symptom, &tmp.Diagnosis, &tmp.Therapy, &tmp.Result, &tmp.UpdatedAt); err != nil {
			log.Printf("[MedicalRecordSelector-FindByEmail] scan rows error: %v", err)
			continue
		}

		result = append(result, &tmp)
	}
	return result, nil
}
