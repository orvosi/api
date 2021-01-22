package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/internal/http/handler"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type MedicalRecordCreator_Executor struct {
	handler *handler.MedicalRecordCreator
	usecase *mock_usecase.MockCreateMedicalRecord
}

func TestNewMedicalRecordCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of MedicalRecordCreator", func(t *testing.T) {
		exec := createMedicalRecordCreatorExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestMedicalRecordCreator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("can't process invalid medical record request", func(t *testing.T) {
		body, _ := json.Marshal("invalid request body")
		req := httptest.NewRequest(http.MethodPost, "/medical-records", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)

		exec := createMedicalRecordCreatorExecutor(ctrl)
		exec.handler.Create(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		str := fmt.Sprintf("%s\n", `{"errors":[{"code":"02-003","message":"Medical record request is invalid. Please, check the JSON request"}],"meta":null}`)
		assert.Equal(t, str, rec.Body.String())
	})
}

func createMedicalRecordCreatorExecutor(ctrl *gomock.Controller) *MedicalRecordCreator_Executor {
	u := mock_usecase.NewMockCreateMedicalRecord(ctrl)
	h := handler.NewMedicalRecordCreator(u)
	return &MedicalRecordCreator_Executor{
		handler: h,
		usecase: u,
	}
}
