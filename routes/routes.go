package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// Routes - routes
func Routes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return e
}
