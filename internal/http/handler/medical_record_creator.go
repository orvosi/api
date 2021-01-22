package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/middleware"
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

	user, err := extractUserFromRequestContext(ctx.Request().Context())
	if err != nil {
		res := response.NewError(err)
		ctx.JSON(http.StatusInternalServerError, res)
		return err
	}

	record := createMedicalRecordFromRequest(&request, user)
	if err := mrc.creator.Create(ctx.Request().Context(), record); err != nil {
		res := response.NewError(err)
		status := http.StatusInternalServerError
		if err.Code != entity.ErrInternalServer.Code {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, res)
		return err
	}

	ctx.JSON(http.StatusCreated, response.NewSuccess(nil, response.EmptyMeta{}))
	return nil
}

func extractUserFromRequestContext(ctx context.Context) (*entity.User, error) {
	val := ctx.Value(middleware.UserContextKey)
	user, ok := val.(*entity.User)
	if !ok {
		return nil, entity.ErrInternalServer
	}
	return user, nil
}

func createMedicalRecordFromRequest(req *MedicalRecordRequest, user *entity.User) *entity.MedicalRecord {
	return &entity.MedicalRecord{
		User:      user,
		Symptom:   req.Symptom,
		Diagnosis: req.Diagnosis,
		Therapy:   req.Therapy,
		Result:    req.Result,
	}
}
