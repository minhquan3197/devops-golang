package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"project-golang/api"
	"project-golang/internal/seed"
	"project-golang/third_party/mongodb"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	godotenv.Load("../../internal/environments/.env")
	uri = os.Getenv("DATABASE_URI")
	database = os.Getenv("DATABASE_PROJECT")
	port = os.Getenv("APP_PORT")
}

func main() {
	mongodb.ConnectMongoDB(uri, database)
	seed.All()

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
		middleware.CORS(),
	)

	e.GET("/", func(c echo.Context) error {
		r := &Reply{
			Response:  "Server is running",
			Timestamp: time.Now().UTC(),
			Random:    rand.Intn(1000),
		}
		sr, _ := json.Marshal(r)
		return c.String(http.StatusOK, string(sr))
	})

	e.Pre(APIVersion)
	api.Router(e)

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
