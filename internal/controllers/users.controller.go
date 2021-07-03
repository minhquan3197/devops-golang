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
	g.GET("/users", list, middlewares.AuthMiddleware)
	g.POST("/users", create, middlewares.AuthMiddleware)
	g.GET("/users/:id", detail, middlewares.AuthMiddleware)
	g.PUT("/users/:id", update, middlewares.AuthMiddleware)
	g.DELETE("/users/:id", remove, middlewares.AuthMiddleware)
}

// @Summary Get all users
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Param search query string false "Search name"
// @Success 200 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /users [get]
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

// @Summary Create user
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param create body interfaces.CreateUser true "Create user"
// @Success 201 {object} response.EchoR
// @Failure 422 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /users [post]
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

// @Summary Get detail user
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path string true "ID User"
// @Success 200 {object} response.EchoR
// @Failure 404 {object} response.EchoR
// @Router /users/{id} [get]
func detail(c echo.Context) error {
	r := response.EchoResponse(c)
	res, err := services.FindUserByID(c.Param("id"))
	if err != nil {
		return r.NotFound()
	}
	return r.OK(res)
}

// @Summary Update user
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param update body interfaces.UpdateUser true "Update user"
// @Param id path string true "ID User"
// @Success 200 {object} response.EchoR
// @Failure 422 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /users/{id} [put]
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
		return r.NotFound()
	}
	err = services.UpdateUserByID(user.ID, req)
	if err != nil {
		return r.BadRequest()
	}
	return r.OK(nil)
}

// @Summary Delete user
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path string true "ID User"
// @Success 200 {object} response.EchoR
// @Failure 404 {object} response.EchoR
// @Router /users/{id} [delete]
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
