package handler

import (
	"net/http"
	"time"

	"github.com/indrasaputra/hashids"
	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/response"
	"github.com/orvosi/api/usecase"
)

// MedicalRecordResponse defines the JSON response of medical record.
type MedicalRecordResponse struct {
	ID        hashids.ID `json:"id"`
	Symptom   string     `json:"symptom"`
	Diagnosis string     `json:"diagnosis"`
	Therapy   string     `json:"therapy"`
	Result    string     `json:"result"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy string     `json:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at"`
}

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

// FindByID handles `GET /medical-records/:id` endpoint.
// It extracts the user's email from bearer token
// then finds all medical records bounded to the user.
func (mf *MedicalRecordFinder) FindByID(ctx echo.Context) error {
	str := ctx.Param("id")
	id, herr := hashids.DecodeHash([]byte(str))
	if herr != nil {
		res := response.NewError(entity.ErrInvalidID)
		ctx.JSON(http.StatusInternalServerError, res)
		return herr
	}

	user, cerr := extractUserFromRequestContext(ctx.Request().Context())
	if cerr != nil {
		res := response.NewError(cerr)
		ctx.JSON(http.StatusInternalServerError, res)
		return cerr
	}

	record, ferr := mf.finder.FindByID(ctx.Request().Context(), uint64(id), user.Email)
	if ferr != nil {
		res := response.NewError(ferr)
		status := http.StatusInternalServerError
		if ferr.Code == entity.ErrUnauthorized.Code {
			status = http.StatusUnauthorized
		}
		ctx.JSON(status, res)
		return ferr
	}

	res := createMedicalRecordResponse(record)
	ctx.JSON(http.StatusOK, response.NewSuccess(res, response.EmptyMeta{}))
	return nil
}

// FindByEmail handles `GET /medical-records` endpoint.
// It extracts the user's email from bearer token
// then finds all medical records bounded to the user.
func (mf *MedicalRecordFinder) FindByEmail(ctx echo.Context) error {
	user, cerr := extractUserFromRequestContext(ctx.Request().Context())
	if cerr != nil {
		res := response.NewError(cerr)
		ctx.JSON(http.StatusInternalServerError, res)
		return cerr
	}

	records, ferr := mf.finder.FindByEmail(ctx.Request().Context(), user.Email)
	if ferr != nil {
		res := response.NewError(ferr)
		status := http.StatusInternalServerError
		if ferr.Code != entity.ErrInternalServer.Code {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, res)
		return ferr
	}

	res := createMedicalRecordResponses(records)
	ctx.JSON(http.StatusOK, response.NewSuccess(res, response.EmptyMeta{}))
	return nil
}

func createMedicalRecordResponses(mrs []*entity.MedicalRecord) []*MedicalRecordResponse {
	var res []*MedicalRecordResponse
	for _, mr := range mrs {
		res = append(res, createMedicalRecordResponse(mr))
	}
	return res
}

func createMedicalRecordResponse(mr *entity.MedicalRecord) *MedicalRecordResponse {
	return &MedicalRecordResponse{
		ID:        mr.ID,
		Symptom:   mr.Symptom,
		Diagnosis: mr.Diagnosis,
		Therapy:   mr.Therapy,
		Result:    mr.Result,
		CreatedBy: mr.CreatedBy,
		CreatedAt: mr.CreatedAt,
		UpdatedBy: mr.UpdatedBy,
		UpdatedAt: mr.UpdatedAt,
	}
}
