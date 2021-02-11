package router

import "github.com/labstack/echo/v4"

// Route defines an HTTP route.
type Route struct {
	// Method defines the HTTP method.
	Method string
	// Path defines the HTTP path.
	Path string
	// Handler defines the handler for the route.
	Handler echo.HandlerFunc
	// Middlewares defines the list of middleware used for the route.
	Middlewares []echo.MiddlewareFunc
}
