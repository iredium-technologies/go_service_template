package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloController -
type HelloController struct {
	Controller
}

func (controller HelloController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
