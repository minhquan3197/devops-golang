package services

import (
	"net/http"
	"project-golang/internal/interfaces"
	"project-golang/pkg/bcrypt"
	"project-golang/pkg/constants"
	"project-golang/pkg/jwt"

	"github.com/labstack/echo/v4"
)

type (
	requestLogin = interfaces.RequestLogin
)

// Login func login for user
func Login(req requestLogin) (string, error) {
	var token string
	res, err := FindUserByUsername(req.Username)
	if err != nil {
		return token, err
	}
	if !bcrypt.ComparePassword(req.Password, res.Password) {
		return token, echo.NewHTTPError(http.StatusUnauthorized, constants.Unauthorized)
	}
	token = jwt.Encrypt("token", res.Username)
	return token, nil
}
