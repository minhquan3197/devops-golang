package users

import (
	"net/http"
	"project-golang/internal/middlewares"

	"project-golang/internal/public/convert"
	"project-golang/internal/public/response"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Router func
func Router(g *echo.Group) {
	g.GET("/users", list, middlewares.AuthMiddleware)
	g.POST("/users", create, middlewares.AuthMiddleware)
	g.GET("/users/:id", detail, middlewares.AuthMiddleware)
	g.PUT("/users/:id", update, middlewares.AuthMiddleware)
	g.DELETE("/users/:id", remove, middlewares.AuthMiddleware)
	g.GET("/users/info", info, middlewares.AuthMiddleware)
}

func list(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	search := c.QueryParam("search")
	page64, err := convert.StringToInt64(page)
	limit64, err := convert.StringToInt64(limit)
	res, err := Paginate(limit64, page64, search)
	if err != nil {
		return response.RespData(c, http.StatusBadRequest, err.Error())
	}
	return response.RespData(c, http.StatusOK, res)
}

func create(c echo.Context) error {
	var req createUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return response.RespData(c, http.StatusUnprocessableEntity, err.Error())
	}
	err = Create(req)
	if err != nil {
		return response.RespData(c, http.StatusBadRequest, err.Error())
	}
	return response.RespData(c, http.StatusOK, nil)
}

func detail(c echo.Context) error {
	id := c.Param("id")
	objectID, err := convert.StringToObjectID(id)
	res, err := FindOne(bson.M{"_id": objectID})
	if err != nil {
		return response.RespData(c, http.StatusBadRequest, err.Error())
	}
	return response.RespData(c, http.StatusOK, res)
}

func update(c echo.Context) error {
	id := c.Param("id")
	var req updateUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return response.RespData(c, http.StatusUnprocessableEntity, err.Error())
	}
	objectID, err := convert.StringToObjectID(id)
	err = Update(objectID, req)
	if err != nil {
		return response.RespData(c, http.StatusBadRequest, err.Error())
	}
	return response.RespData(c, http.StatusOK, nil)
}

func remove(c echo.Context) error {
	id := c.Param("id")
	objectID, _ := convert.StringToObjectID(id)
	Remove(bson.M{"_id": objectID})
	return response.RespData(c, http.StatusOK, nil)
}

func info(c echo.Context) error {
	user := c.Get("user")
	return response.RespData(c, http.StatusOK, user)
}
