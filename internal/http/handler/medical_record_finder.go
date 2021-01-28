package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/response"
	"github.com/orvosi/api/usecase"
)

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

// FindByEmail handles `GET /medical-records` endpoint.
// It extracts the user's email from bearer token
// then finds all medical records bounded to the user.
func (mf *MedicalRecordFinder) FindByEmail(ctx echo.Context) error {
	user, err := extractUserFromRequestContext(ctx.Request().Context())
	if err != nil {
		res := response.NewError(err)
		ctx.JSON(http.StatusInternalServerError, res)
		return err
	}

	res, ferr := mf.finder.FindByEmail(ctx.Request().Context(), user.Email)
	if ferr != nil {
		status := http.StatusInternalServerError
		if ferr.Code != entity.ErrInternalServer.Code {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, res)
		return ferr
	}

	ctx.JSON(http.StatusOK, response.NewSuccess(res, response.EmptyMeta{}))
	return nil
}
