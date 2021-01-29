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

// UpdateMedicalRecordStorage defines the business logic
// to update a medical record into a storage.
type UpdateMedicalRecordStorage interface {
	// FindByID finds medical records by its id.
	FindByIDAndEmail(ctx context.Context, id uint64, email string) (*entity.MedicalRecord, *entity.Error)
	// Update updates certain medical record.
	Update(ctx context.Context, id uint64, record *entity.MedicalRecord) *entity.Error
}

// MedicalRecordUpdater responsibles for medical record update workflow.
type MedicalRecordUpdater struct {
}

// NewMedicalRecordUpdater creates an instance of MedicalRecordUpdater.
func NewMedicalRecordUpdater() *MedicalRecordUpdater {
	return &MedicalRecordUpdater{}
}

// Update updates the medical record.
func (mu *MedicalRecordUpdater) Update(ctx context.Context, record *entity.MedicalRecord, email string) *entity.Error {
	return nil
}
