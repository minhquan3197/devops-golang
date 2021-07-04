package response

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type EchoR struct {
	c          echo.Context
	StatusCode int         `json:"-"`
	Code       int64       `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data,omitempty"`
}

func EchoResponse(c echo.Context) EchoR {
	return EchoR{c: c}
}

func (r EchoR) Unauthorized() error {
	r.Code = 401
	r.Msg = "Unauthorized"
	return r.c.JSON(http.StatusUnauthorized, r)
}

func (r EchoR) BadRequest(err string) error {
	r.Code = 400
	if err == "" {
		r.Msg = "Bad Request"
	} else {
		r.Msg = err
	}
	return r.c.JSON(http.StatusBadRequest, r)
}

func (r EchoR) NotFound() error {
	r.Code = 404
	r.Msg = "Not Found"
	return r.c.JSON(http.StatusNotFound, r)
}

func (r EchoR) UnprocessableEntity(err error) error {
	validateErrors := strings.Split(err.Error(), ";")
	r.Code = 422
	r.Msg = validateErrors[0]
	r.Data = validateErrors
	return r.c.JSON(http.StatusUnprocessableEntity, r)
}

func (r EchoR) Err(err error) error {
	r.Code = 400
	r.Msg = err.Error()
	return r.c.JSON(http.StatusBadRequest, r)
}

func (r EchoR) OK(data interface{}) error {
	r.Code = 200
	r.Msg = "Success"
	r.Data = data
	return r.c.JSON(http.StatusOK, r)
}

func (r EchoR) Created() error {
	r.Code = 201
	r.Msg = "Created"
	return r.c.JSON(http.StatusOK, r)
}

func (r EchoR) NoContet() error {
	r.Code = 204
	r.Msg = "No content"
	return r.c.JSON(http.StatusOK, r)
}
