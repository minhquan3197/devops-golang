package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"project-golang/api/routers"
	"project-golang/configs"
	"project-golang/migrations"
	"project-golang/pkg/mongodb"
	"time"

	_ "project-golang/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Reply Log
type Reply struct {
	Response  string    `json:"response"`
	Timestamp time.Time `json:"timestamp"`
	Random    int       `json:"random"`
}

var (
	uri      string
	database string
	port     string
)

func init() {
	env := os.Getenv("GOLANG_ENV")
	fmt.Println(env)
	godotenv.Load()
	cfg := configs.Load()
	uri = cfg.DatabaseUri
	database = cfg.DatabaseProject
	port = cfg.Port
}

func healCheck(c echo.Context) error {
	r := &Reply{
		Response:  "Server is running",
		Timestamp: time.Now().UTC(),
		Random:    rand.Intn(1000),
	}
	sr, _ := json.Marshal(r)
	return c.String(http.StatusOK, string(sr))
}

func Excute() {
	mongodb.ConnectMongoDB(uri, database)
	migrations.All()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
		middleware.CORS(),
	)

	e.GET("/", healCheck)

	e.Pre(APIVersion)
	routers.Router(e)

	e.Logger.Fatal(e.Start(":" + port))
}

// APIVersion Header Based Versioning
func APIVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		apiVersion := headers.Get("version")
		if apiVersion != "" {
			req.URL.Path = fmt.Sprintf("/%s%s", apiVersion, req.URL.Path)
			return next(c)
		}
		return next(c)
	}
}
