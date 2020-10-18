package auth

import (
	"net/http"
	"project-golang/internal/public/response"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// Router func
func Router(g *echo.Group) {
	g.POST("/auth/login", login)
}

func login(c echo.Context) error {
	var req requestLogin
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return response.RespData(c, http.StatusUnprocessableEntity, err.Error())
	}
	res, err := Login(req)
	if err != nil {
		return response.RespData(c, http.StatusUnauthorized, nil)
	}
	return response.RespData(c, http.StatusOK, res)
}
