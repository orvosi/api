package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/hashids"
	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/handler"
	"github.com/orvosi/api/internal/http/middleware"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordFinder_Executor struct {
	handler *handler.MedicalRecordFinder
	usecase *mock_usecase.MockFindMedicalRecord
}

func TestNewMedicalRecordFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordFinder", func(t *testing.T) {
		exec := createMedicalRecordFinderExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestMedicalRecordFinder_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("can't extract user information from request context", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/medical-records", nil)

		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)

		exec := createMedicalRecordFinderExecutor(ctrl)
		exec.handler.FindByEmail(ctx)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		str := fmt.Sprintf("%s\n", `{"errors":[{"code":"01-001","message":"Internal server error"}],"meta":null}`)
		assert.Equal(t, str, rec.Body.String())
	})

	t.Run("finder service returns 4xx error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/medical-records", nil)
		user := createUserInformation()
		req = req.WithContext(context.WithValue(context.Background(), middleware.ContextKeyUser, user))

		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)

		exec := createMedicalRecordFinderExecutor(ctrl)
		exec.usecase.EXPECT().FindByEmail(ctx.Request().Context(), user.Email).Return([]*entity.MedicalRecord{}, entity.ErrInvalidEmail)
		exec.handler.FindByEmail(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		str := fmt.Sprintf("%s\n", `{"errors":[{"code":"02-004","message":"Email is invalid. Please, check the email"}],"meta":null}`)
		assert.Equal(t, str, rec.Body.String())
	})

	t.Run("finder service returns 5xx error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/medical-records", nil)
		user := createUserInformation()
		req = req.WithContext(context.WithValue(context.Background(), middleware.ContextKeyUser, user))

		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)

		exec := createMedicalRecordFinderExecutor(ctrl)
		exec.usecase.EXPECT().FindByEmail(ctx.Request().Context(), user.Email).Return([]*entity.MedicalRecord{}, entity.ErrInternalServer)
		exec.handler.FindByEmail(ctx)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		str := fmt.Sprintf("%s\n", `{"errors":[{"code":"01-001","message":"Internal server error"}],"meta":null}`)
		assert.Equal(t, str, rec.Body.String())
	})

	t.Run("successfully get medical records from certain user (email)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/medical-records", nil)
		user := createUserInformation()
		req = req.WithContext(context.WithValue(context.Background(), middleware.ContextKeyUser, user))

		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)

		exec := createMedicalRecordFinderExecutor(ctrl)
		mrs := createMedicalRecords()
		exec.usecase.EXPECT().FindByEmail(ctx.Request().Context(), user.Email).Return(mrs, nil)
		exec.handler.FindByEmail(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		str := fmt.Sprintf("%s\n", `{"data":[{"id":"oWx0b8DZ1a","user":{"id":"oWx0b8DZ1a","email":"user@dummy.com","name":"User Dummy","google_id":"1234567890","created_at":"","updated_at":"","created_by":"0001-01-01T00:00:00Z","updated_by":"0001-01-01T00:00:00Z"},"symptom":"Symptom","diagnosis":"Diagnosis","therapy":"Therapy","result":"Result","created_at":"user@dummy.com","updated_at":"user@dummy.com","created_by":"2021-01-28T15:00:00Z","updated_by":"2021-01-28T15:00:00Z"}],"meta":{}}`)
		assert.Equal(t, str, rec.Body.String())
	})
}

func createMedicalRecords() []*entity.MedicalRecord {
	return []*entity.MedicalRecord{
		&entity.MedicalRecord{
			ID: hashids.ID(1),
			User: &entity.User{
				ID:       hashids.ID(1),
				Email:    "user@dummy.com",
				Name:     "User Dummy",
				GoogleID: "1234567890",
			},
			Symptom:   "Symptom",
			Diagnosis: "Diagnosis",
			Therapy:   "Therapy",
			Result:    "Result",
			Auditable: entity.Auditable{
				CreatedBy: "user@dummy.com",
				CreatedAt: time.Date(2021, time.January, 28, 15, 00, 00, 00, time.UTC),
				UpdatedBy: "user@dummy.com",
				UpdatedAt: time.Date(2021, time.January, 28, 15, 00, 00, 00, time.UTC),
			},
		},
	}
}

func createMedicalRecordFinderExecutor(ctrl *gomock.Controller) *MedicalRecordFinder_Executor {
	u := mock_usecase.NewMockFindMedicalRecord(ctrl)
	h := handler.NewMedicalRecordFinder(u)
	return &MedicalRecordFinder_Executor{
		handler: h,
		usecase: u,
	}
}
