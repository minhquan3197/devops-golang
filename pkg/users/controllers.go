package users

import (
	"project-golang/internal/middlewares"

	"project-golang/internal/response"
	"project-golang/utils/convert"

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
	r := response.EchoResponse(c)
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	search := c.QueryParam("search")
	page64, err := convert.StringToInt64(page)
	limit64, err := convert.StringToInt64(limit)
	res, err := Paginate(limit64, page64, search)

	if err != nil {
		return r.BadRequest()
	}
	return r.OK(res)
}

func create(c echo.Context) error {
	r := response.EchoResponse(c)
	var req CreateUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	err = Create(req)
	if err != nil {
		return r.BadRequest()
	}
	return r.Created()
}

func detail(c echo.Context) error {
	r := response.EchoResponse(c)
	id := c.Param("id")
	objectID, err := convert.StringToObjectID(id)
	res, err := FindOne(bson.M{"_id": objectID})
	if err != nil {
		return r.BadRequest()
	}
	return r.OK(res)
}

func update(c echo.Context) error {
	r := response.EchoResponse(c)
	id := c.Param("id")
	var req UpdateUser
	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return r.UnprocessableEntity(err)
	}
	objectID, err := convert.StringToObjectID(id)
	err = Update(objectID, req)
	if err != nil {
		return r.BadRequest()
	}
	return r.OK(nil)
}

func remove(c echo.Context) error {
	r := response.EchoResponse(c)
	id := c.Param("id")
	objectID, _ := convert.StringToObjectID(id)
	Remove(bson.M{"_id": objectID})
	return r.NoContet()
}

func info(c echo.Context) error {
	r := response.EchoResponse(c)
	user := c.Get("user")
	return r.OK(user)
}
