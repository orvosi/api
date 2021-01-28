package handler

import "github.com/orvosi/api/usecase"

// MedicalRecordFinder handles HTTP request and response
// for find medical record.
type MedicalRecordFinder struct {
	finder usecase.FindMedicalRecord
}

// NewMedicalRecordFinder creates an instance of MedicalFinder.
func NewMedicalRecordFinder(finder usecase.FindMedicalRecord) *MedicalRecordFinder {
	return &MedicalRecordFinder{
		finder: finder,
	}
}
