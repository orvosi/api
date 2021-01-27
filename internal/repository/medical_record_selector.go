package repository

import "database/sql"

// MedicalRecordSelector connects the database with medical record entity
// and only responsible for retrieving medical record data.
type MedicalRecordSelector struct {
	db *sql.DB
}

// NewMedicalRecordSelector creates an instance of MedicalRecordSelector.
func NewMedicalRecordSelector(db *sql.DB) *MedicalRecordSelector {
	return &MedicalRecordSelector{db: db}
}
