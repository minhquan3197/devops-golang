package auth

import (
	"project-golang/internal/response"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// Router func
func Router(g *echo.Group) {
	g.POST("/auth/login", login)
	g.GET("/auth/info", info)
}

func login(c echo.Context) error {
	var req RequestLogin
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	r := response.EchoResponse(c)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	res, err := Login(req)
	if err != nil {
		return r.Unauthorized()
	}
	return r.OK(res)
}

func info(c echo.Context) error {
	r := response.EchoResponse(c)
	return r.OK(nil)
}
