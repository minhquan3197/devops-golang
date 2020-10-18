package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// RespStruct struct define resposne data
type RespStruct struct {
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// RespData this is function
func RespData(c echo.Context, statusCode int, responseData interface{}) error {
	switch statusCode {
	case http.StatusOK:
		return c.JSON(http.StatusOK, RespStruct{
			Code:    statusCode,
			Data:    responseData,
			Message: "OK",
			Status:  true,
		})
	case http.StatusForbidden:
		return c.JSON(http.StatusForbidden, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: "Forbidden",
			Status:  false,
		})
	case http.StatusUnauthorized:
		return c.JSON(http.StatusUnauthorized, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: "Unauthentication",
			Status:  false,
		})
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: responseData.(string),
			Status:  false,
		})
	case http.StatusUnprocessableEntity:
		fmt.Println(responseData)
		stringArrayData := strings.Split(responseData.(string), ";")
		fmt.Println(stringArrayData)
		return c.JSON(http.StatusUnprocessableEntity, RespStruct{
			Code:    statusCode,
			Data:    stringArrayData,
			Message: stringArrayData[0],
			Status:  false,
		})
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: "Not found",
			Status:  false,
		})
	case http.StatusConflict:
		return c.JSON(http.StatusConflict, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: responseData.(string),
			Status:  false,
		})
	default:
		return c.JSON(http.StatusInternalServerError, RespStruct{
			Code:    statusCode,
			Data:    nil,
			Message: "Server error",
			Status:  false,
		})
	}
}
