package usecase

import (
	"context"

	"github.com/orvosi/api/entity"
)

// CreateMedicalRecord defines the business logic
// to create a medical record.
type CreateMedicalRecord interface {
	// Create creates a new medical record.
	Create(ctx context.Context, record *entity.MedicalRecord) *entity.Error
}
