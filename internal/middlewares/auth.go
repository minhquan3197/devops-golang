package middlewares

import (
	"net/http"
	"project-golang/internal/private/jwt"
	"project-golang/internal/public/response"
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

		hToken := c.Request().Header.Get("Authorization")
		if hToken == "" {
			return response.RespData(c, http.StatusUnauthorized, nil)
		}
		token := strings.Split(hToken, " ")[1]
		results := jwt.Decrypt(token)
		if results == nil {
			return response.RespData(c, http.StatusUnauthorized, nil)
		}

		result, err := authService.FindUserWithUsername(results["token"].(string))

		if err != nil {
			return response.RespData(c, http.StatusUnauthorized, nil)
		}

		// Remove password
		result.Password = ""
		c.Set("user", result)
		return next(c)
	}
}
