package routers

import (
	"project-golang/internal/controllers"

	"github.com/labstack/echo/v4"
)

// Router func
func Router(e *echo.Echo) {

	api := e.Group("/api/v1")

	controllers.RouterAuth(api)
	controllers.RouterUser(api)
}
