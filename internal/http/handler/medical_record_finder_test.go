package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
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
}

func createMedicalRecordFinderExecutor(ctrl *gomock.Controller) *MedicalRecordFinder_Executor {
	u := mock_usecase.NewMockFindMedicalRecord(ctrl)
	h := handler.NewMedicalRecordFinder(u)
	return &MedicalRecordFinder_Executor{
		handler: h,
		usecase: u,
	}
}
