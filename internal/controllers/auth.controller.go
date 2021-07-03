package controllers

import (
	"fmt"
	"project-golang/api/middlewares"
	"project-golang/api/response"
	"project-golang/internal/interfaces"
	"project-golang/internal/services"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// Router func
func RouterAuth(g *echo.Group) {
	g.POST("/auth/login", login)
	g.GET("/auth/info", info, middlewares.AuthMiddleware)
}

// Login user
// @Summary Login user.
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param login body interfaces.RequestLogin true "Login payload"
// @Success 200 {object} response.EchoR
// @Failure 401 {object} response.EchoR
// @Router /auth/login [post]
func login(c echo.Context) error {
	var req interfaces.RequestLogin
	c.Bind(&req)
	fmt.Println(req)
	_, err := govalidator.ValidateStruct(req)
	r := response.EchoResponse(c)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	res, err := services.Login(req)
	if err != nil {
		return r.Unauthorized()
	}
	return r.OK(res)
}

// Info user
// @Summary Get info user by token.
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.EchoR
// @Failure 401 {object} response.EchoR
// @Router /auth/info [get]
func info(c echo.Context) error {
	r := response.EchoResponse(c)
	user := c.Get("user")
	return r.OK(user)
}
