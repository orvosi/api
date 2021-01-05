package repository

import "database/sql"

// MedicalRecordInserter connects the database with medical record entity
// and only responsible for inserting a new data.
type MedicalRecordInserter struct {
	db *sql.DB
}

// NewMedicalRecordInserter creates an instance of MedicalRecordInserter.
func NewMedicalRecordInserter(db *sql.DB) *MedicalRecordInserter {
	return &MedicalRecordInserter{db: db}
}
