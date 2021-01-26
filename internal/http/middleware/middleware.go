package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/orvosi/api/entity"
	"github.com/orvosi/api/internal/http/response"
)

// ContextKey is just an alias for string to be used
// as key when assign a value in context.
// This is to avoid go-lint warning
// "should not use basic type untyped string as key in context.WithValue".
type ContextKey string

const (
	// ContextKeyUser is just a string "user" defined as a key
	// to save a user information in context.
	ContextKeyUser = ContextKey("user")

	authHeaderKey = "Authorization"
	authBearerKey = "Bearer"
)

// JWTDecoder defines the function contract to decode JWT.
type JWTDecoder func(token string) (*entity.User, *entity.Error)

// WithJWTDecoder decodes the Bearer authorization that contains JWT
// into a detailed user information.
// Then, the user information will be passed through the request context.
func WithJWTDecoder(decoder JWTDecoder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			val := ctx.Request().Header.Get(authHeaderKey)
			token := strings.Split(val, " ")
			if len(token) != 2 || (len(token) == 2 && token[0] != authBearerKey) {
				res := response.NewError(entity.ErrUnauthorized)
				ctx.JSON(http.StatusUnauthorized, res)
				return entity.ErrUnauthorized
			}

			user, err := decoder(token[1])
			if err != nil {
				res := response.NewError(entity.ErrUnauthorized)
				ctx.JSON(http.StatusUnauthorized, res)
				return err
			}

			reqCtx := context.WithValue(ctx.Request().Context(), ContextKeyUser, user)
			req := ctx.Request().Clone(reqCtx)
			ctx.SetRequest(req)

			return next(ctx)
		}
	}
}
