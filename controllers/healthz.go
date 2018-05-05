package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthzController - health checking controller
type HealthzController struct {
	Controller
}

// Index - should return OK
func (controller HealthzController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
