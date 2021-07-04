package controllers

import (
	"fmt"
	"project-golang/api/response"
	"project-golang/configs"
	"strconv"
	"strings"
	"time"

	awsPkg "project-golang/pkg/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	cfg := configs.Load()
	r := response.EchoResponse(c)
	sess := awsPkg.ConnectAwsService()
	uploader := s3manager.NewUploader(sess)
	file, _ := c.FormFile("upload")
	readFile, _ := file.Open()
	defer readFile.Close()
	spl := strings.Split(file.Filename, ".")
	uploadedFileName := strings.Join(spl, ".")
	t := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(cfg.AWS.Bucket)
	keyFile := "haven-gate/RosTest/" + t + "-" + uploadedFileName
	fmt.Println(keyFile)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.AWS.Bucket),
		ACL:    aws.String(cfg.AWS.ACL),
		Key:    aws.String(keyFile),
		Body:   readFile,
	})
	fmt.Println(readFile)
	if err != nil {
		return r.BadRequest(err.Error())
	}
	fmt.Println(res)
	return r.OK(res)
}
