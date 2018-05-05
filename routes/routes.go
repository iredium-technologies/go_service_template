package routes

import (
	"github.com/iredium-technologies/go_service_template/controllers"
	"github.com/labstack/echo"
)

// Routes - routes
func Routes(e *echo.Echo) *echo.Echo {
	helloController := controllers.HelloController{}
	healthzController := controllers.HealthzController{}

	e.GET("/", helloController.Index)
	e.GET("/healthz", healthzController.Index)

	return e
}
