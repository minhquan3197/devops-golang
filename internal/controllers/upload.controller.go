package controllers

import (
	"fmt"
	"project-golang/api/response"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// Router func
func RouterUpload(g *echo.Group) {
	g.POST("/upload", upload)
}

// @Summary Upload file
// @Tags Upload
// @Accept multipart/form-data
// @Produce application/json
// @Param upload formData file true "File payload"
// @Success 200 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /upload [post]
func upload(c echo.Context) error {
	r := response.EchoResponse(c)
	file, _ := c.FormFile("upload")
	readFile, _ := file.Open()
	defer readFile.Close()
	spl := strings.Split(file.Filename, ".")
	uploadedFileName := strings.Join(spl, ".")
	t := strconv.FormatInt(time.Now().Unix(), 10)
	keyFile := t + "-" + uploadedFileName
	fmt.Println(keyFile)
	return r.OK(file)
}
