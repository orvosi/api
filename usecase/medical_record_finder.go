package usecase

import (
	"context"
	"net"
	"regexp"
	"strings"

	"github.com/orvosi/api/entity"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// FindMedicalRecord defines the business logic
// to find a medical record.
type FindMedicalRecord interface {
	// FindByID finds all medical records by its ID and validated by user's email.
	// Only medical record that is owned by the right user's email can be retrieved.
	// If the medical record is not owned by the user, it will return ErrUnauthorized.
	FindByID(ctx context.Context, id uint64, email string) (*entity.MedicalRecord, *entity.Error)
	// FindByEmail finds medical records that belong to specific user (based on email).
	FindByEmail(ctx context.Context, email string) ([]*entity.MedicalRecord, *entity.Error)
}

// MedicalRecordSelector defines the business logic
// to select or find medical record data from storage.
type MedicalRecordSelector interface {
	// FindByID finds medical records by its id.
	FindByID(ctx context.Context, id uint64) (*entity.MedicalRecord, *entity.Error)
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

// FindByEmail finds medical records that belong to specific user (based on email).
// The email will be verified first using regex and LookupMX.
func (mf *MedicalRecordFinder) FindByEmail(ctx context.Context, email string) ([]*entity.MedicalRecord, *entity.Error) {
	if err := validateEmail(email); err != nil {
		return []*entity.MedicalRecord{}, entity.ErrInvalidEmail
	}

	return mf.selector.FindByEmail(ctx, email)
}

func validateEmail(email string) *entity.Error {
	if !emailRegex.MatchString(email) {
		return entity.ErrInvalidEmail
	}

	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return entity.ErrInvalidEmail
	}
	return nil
}
