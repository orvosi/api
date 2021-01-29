package usecase

import (
	"context"

	"github.com/orvosi/api/entity"
)

// UpdateMedicalRecord defines the business logic
// to update a medical record.
type UpdateMedicalRecord interface {
	// Update updates a new medical record.
	Update(ctx context.Context, record *entity.MedicalRecord) *entity.Error
}

// MedicalRecordUpdater responsibles for medical record update workflow.
type MedicalRecordUpdater struct {
}

// NewMedicalRecordUpdater creates an instance of MedicalRecordUpdater.
func NewMedicalRecordUpdater() *MedicalRecordUpdater {
	return &MedicalRecordUpdater{}
}
