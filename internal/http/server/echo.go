package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/orvosi/api/internal/http/router"
)

// Server acts as echo.Echo server.
type Server struct {
	*echo.Echo
}

// NewServer creates an instance of Echo.
func NewServer(routes []*router.Route) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	for _, route := range routes {
		e.Add(route.Method, route.Path, route.Handler)
	}

	return &Server{e}
}
