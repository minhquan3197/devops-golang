package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Reply Log
type Reply struct {
	Response  string    `json:"response"`
	Timestamp time.Time `json:"timestamp"`
	Random    int       `json:"random"`
	Database  string    `json:"database"`
}

var (
	c  *mongo.Client
	db string
)

func main() {
	e := echo.New()
	godotenv.Load()
	env := os.Getenv("GOLANG_ENV")
	if env == "production" {
		godotenv.Load(".env.production")
	} else {
		godotenv.Load(".env.develop")
	}
	Port := os.Getenv("APP_PORT")
	// URI := os.Getenv("DATABASE_URI")
	URI := "mongodb://mongo:27017/"
	ctx := context.Background()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		db = "Unable to connect to database : " + URI
	}
	if c != nil {
		db = "Database connected : " + URI
	}
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
	)

	e.GET("/", func(c echo.Context) error {
		r := &Reply{
			Response:  "Server is running",
			Timestamp: time.Now().UTC(),
			Random:    rand.Intn(1000),
			Database:  db,
		}
		sr, _ := json.Marshal(r)
		return c.String(http.StatusOK, string(sr))
	})

	e.Pre(APIVersion)

	e.Logger.Fatal(e.Start(":" + Port))
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
