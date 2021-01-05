package usecase

import (
	"context"
	"strings"

	"github.com/orvosi/api/entity"
)

// CreateMedicalRecord defines the business logic
// to create a medical record.
type CreateMedicalRecord interface {
	// Create creates a new medical record.
	Create(ctx context.Context, record *entity.MedicalRecord) *entity.Error
}

// MedicalRecordCreator responsibles for medical record creation workflow.
type MedicalRecordCreator struct {
}

// Create creates a new medical record and persist it into a storage.s
func (mrc *MedicalRecordCreator) Create(ctx context.Context, record *entity.MedicalRecord) *entity.Error {
	if err := validateMedicalRecord(record); err != nil {
		return err
	}
	return nil
}

func validateMedicalRecord(record *entity.MedicalRecord) *entity.Error {
	if record == nil {
		return entity.ErrEmptyMedicalRecord
	}

	sanitizeMedicalRecord(record)
	return nil
}

func sanitizeMedicalRecord(record *entity.MedicalRecord) {
	record.Symptom = strings.TrimSpace(record.Symptom)
	record.Diagnosis = strings.TrimSpace(record.Diagnosis)
	record.Therapy = strings.TrimSpace(record.Therapy)
}
