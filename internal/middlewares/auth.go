package middlewares

import (
	"project-golang/internal/private/jwt"
	"project-golang/internal/response"
	authService "project-golang/pkg/auth"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware func
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Connect S3, if use, open it
		// awsS3 := s3.ConnectAws()
		// c.Set("s3", awsS3)

		r := response.EchoResponse(c)

		hToken := c.Request().Header.Get("Authorization")

		if hToken == "" {
			return r.Unauthorized()
		}
		token := strings.Split(hToken, " ")[1]
		results := jwt.Decrypt(token)
		if results == nil {
			return r.Unauthorized()
		}

		result, err := authService.FindUserWithUsername(results["token"].(string))

		if err != nil {
			return r.Unauthorized()
		}

		// Remove password
		result.Password = ""
		c.Set("user", result)
		return next(c)
	}
}
