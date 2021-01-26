package usecase

import (
	"context"

	"github.com/orvosi/api/entity"
)

// FindMedicalRecord defines the business logic
// to find a medical record.
type FindMedicalRecord interface {
	// Find finds medical records that belong to specific user (based on email).
	FindByEmail(ctx context.Context, email string) ([]*entity.MedicalRecord, *entity.Error)
}

// MedicalRecordSelector defines the business logic
// to select or find medical record data from storage.
type MedicalRecordSelector interface {
	// FindByEmail finds all medical records bounded to specific email.
	FindByEmail(ctx context.Context, email string) ([]*entity.MedicalRecord, *entity.Error)
}

// MedicalRecordFinder responsibles for medical record find workflow.
type MedicalRecordFinder struct {
	selector MedicalRecordSelector
}

// NewMedicalRecordFinder creates an instance of MedicalRecordFinder.
func NewMedicalRecordFinder(selector MedicalRecordSelector) *MedicalRecordFinder {
	return &MedicalRecordFinder{
		selector: selector,
	}
}
