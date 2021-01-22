package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/response"
	"github.com/orvosi/api/usecase"
)

// MedicalRecordRequest represents medical record request.
type MedicalRecordRequest struct {
	Symptom   string `json:"symptom"`
	Diagnosis string `json:"diagnosis"`
	Therapy   string `json:"therapy"`
	Result    string `json:"result"`
}

// MedicalRecordCreator handles HTTP request and response
// for create medical record.
type MedicalRecordCreator struct {
	creator usecase.MedicalRecordCreator
}

// NewMedicalRecordCreator creates an instance of MedicalCreator.
func NewMedicalRecordCreator(creator usecase.MedicalRecordCreator) *MedicalRecordCreator {
	return &MedicalRecordCreator{
		creator: creator,
	}
}

// Create handles `POST /medical-records` endpoint.
func (mrc *MedicalRecordCreator) Create(ctx echo.Context) error {
	var request MedicalRecordRequest
	if err := ctx.Bind(&request); err != nil {
		res := response.NewError(entity.ErrInvalidMedicalRecordRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return err
	}

	return nil
}
