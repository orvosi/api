package repository

import "database/sql"

// MedicalRecordUpdater connects the database with medical record entity
// and only responsible for usecase of updating a data.
type MedicalRecordUpdater struct {
	db *sql.DB
}

// NewMedicalRecordUpdater creates an instance of MedicalRecordUpdater.
func NewMedicalRecordUpdater(db *sql.DB) *MedicalRecordUpdater {
	return &MedicalRecordUpdater{db: db}
}
