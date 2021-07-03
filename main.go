package main

import (
	"project-golang/cmd/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8200
// @BasePath /api/v1
func main() {
	api.Excute()
}
