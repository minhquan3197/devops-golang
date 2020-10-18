package api

import (
	auth "project-golang/pkg/auth"
	users "project-golang/pkg/users"

	"github.com/labstack/echo/v4"
)

// Router func
func Router(e *echo.Echo) {

	api := e.Group("/api/v1")

	auth.Router(api)
	users.Router(api)
}
