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
