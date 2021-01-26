package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/middleware"
	"github.com/stretchr/testify/assert"
)

func TestWithJWTDecoder(t *testing.T) {
	t.Run("request doesn't contain authorization", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, rec)

		hdr := createHandler()
		dec := createNormalDecoder()
		hdr = middleware.WithJWTDecoder(dec)(hdr)

		err := hdr(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrUnauthorized, err)
	})

	t.Run("token length is not 2 (invalid)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "invalid-token")
		rec := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, rec)

		hdr := createHandler()
		dec := createNormalDecoder()
		hdr = middleware.WithJWTDecoder(dec)(hdr)

		err := hdr(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrUnauthorized, err)
	})

	t.Run("token prefix is not `Bearer`", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "not-bearer jwt-token")
		rec := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, rec)

		hdr := createHandler()
		dec := createNormalDecoder()
		hdr = middleware.WithJWTDecoder(dec)(hdr)

		err := hdr(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrUnauthorized, err)
	})
}

func createErrorDecoder() middleware.JWTDecoder {
	return func(token string) (*entity.User, *entity.Error) {
		return nil, entity.ErrInternalServer
	}
}

func createNormalDecoder() middleware.JWTDecoder {
	return func(token string) (*entity.User, *entity.Error) {
		return &entity.User{}, nil
	}
}

func createHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
}
