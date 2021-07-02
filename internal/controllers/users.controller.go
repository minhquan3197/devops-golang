package controllers

import (
	"project-golang/api/middlewares"
	"project-golang/api/response"
	"project-golang/internal/interfaces"
	"project-golang/internal/services"
	"project-golang/pkg/convert"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// Router func
func RouterUser(g *echo.Group) {
	g.GET("/users", list)
	g.POST("/users", create, middlewares.AuthMiddleware)
	g.GET("/users/:id", detail, middlewares.AuthMiddleware)
	g.PUT("/users/:id", update, middlewares.AuthMiddleware)
	g.DELETE("/users/:id", remove, middlewares.AuthMiddleware)
}

func list(c echo.Context) error {
	r := response.EchoResponse(c)
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	search := c.QueryParam("search")
	page64, err := convert.StringToInt64(page)
	limit64, err := convert.StringToInt64(limit)
	res, err := services.PaginateUsers(limit64, page64, search)

	if err != nil {
		return r.BadRequest()
	}
	return r.OK(res)
}

func create(c echo.Context) error {
	r := response.EchoResponse(c)
	var req interfaces.CreateUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	err = services.Register(req)
	if err != nil {
		return r.BadRequest()
	}
	return r.Created()
}

func detail(c echo.Context) error {
	r := response.EchoResponse(c)
	res, err := services.FindUserByID(c.Param("id"))
	if err != nil {
		return r.BadRequest()
	}
	return r.OK(res)
}

func update(c echo.Context) error {
	r := response.EchoResponse(c)
	var req interfaces.UpdateUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	user, err := services.FindUserByID(c.Param("id"))
	if err != nil {
		return r.BadRequest()
	}
	err = services.UpdateUserByID(user.ID, req)
	if err != nil {
		return r.BadRequest()
	}
	return r.OK(nil)
}

func remove(c echo.Context) error {
	r := response.EchoResponse(c)
	user, err := services.FindUserByID(c.Param("id"))
	if err != nil {
		return r.BadRequest()
	}
	err = services.RemoveUserByID(user.ID)
	if err != nil {
		return r.BadRequest()
	}
	return r.NoContet()
}
