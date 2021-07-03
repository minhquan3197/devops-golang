package middlewares

import (
	"project-golang/api/response"
	"project-golang/internal/services"
	"project-golang/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware func
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := response.EchoResponse(c)
		hToken := c.Request().Header.Get("Authorization")

		if hToken == "" {
			return r.Unauthorized()
		}
		results := jwt.Decrypt(hToken)
		if results == nil {
			return r.Unauthorized()
		}
		username := results["token"].(string)
		result, err := services.FindUserByUsername(username)

		if err != nil {
			return r.Unauthorized()
		}

		c.Set("user", result)
		return next(c)
	}
}
