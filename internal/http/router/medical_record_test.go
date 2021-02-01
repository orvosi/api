package router_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orvosi/api/internal/http/handler"
	"github.com/orvosi/api/internal/http/router"
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

func TestMedicalRecordCreatorRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("all desired medical record creator routes are registered", func(t *testing.T) {
		desired := map[string]string{
			"/medical-records": "POST",
		}

		h := createMedicalRecordCreator(ctrl)
		routes := router.MedicalRecordCreator(h)

		for _, route := range routes {
			assert.Equal(t, desired[route.Path], route.Method)
		}
	})
}

func TestMedicalRecordFinderRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("all desired medical record finder routes are registered", func(t *testing.T) {
		desired := map[string]string{
			"/medical-records":     "GET",
			"/medical-records/:id": "GET",
		}

		h := createMedicalRecordFinder(ctrl)
		routes := router.MedicalRecordFinder(h)

		for _, route := range routes {
			assert.Equal(t, desired[route.Path], route.Method)
		}
	})
}

func createMedicalRecordCreator(ctrl *gomock.Controller) *handler.MedicalRecordCreator {
	m := mock_usecase.NewMockCreateMedicalRecord(ctrl)
	return handler.NewMedicalRecordCreator(m)
}

func createMedicalRecordFinder(ctrl *gomock.Controller) *handler.MedicalRecordFinder {
	m := mock_usecase.NewMockFindMedicalRecord(ctrl)
	return handler.NewMedicalRecordFinder(m)
}
