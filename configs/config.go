package configs

import (
	"os"
)

const (
	ModeDev        = "dev"
	ModeProduction = "prod"
)

type EnvConfig struct {
	ServerMode      string
	Port            string
	DatabaseUri     string
	DatabaseProject string
	JWTKey          string
}

func Load() EnvConfig {
	cfg := EnvConfig{
		ServerMode:      os.Getenv("SERVER_MODE"),
		Port:            os.Getenv("APP_PORT"),
		DatabaseUri:     os.Getenv("DATABASE_URI"),
		DatabaseProject: os.Getenv("DATABASE_PROJECT"),
		JWTKey:          os.Getenv("JWT_KEY"),
	}

	return cfg
}
