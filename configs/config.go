package configs

import (
	"os"
)

type EnvConfig struct {
	Port            string
	DatabaseUri     string
	DatabaseProject string
	JWTKey          string
}

func Load() EnvConfig {
	cfg := EnvConfig{
		Port:            os.Getenv("APP_PORT"),
		DatabaseUri:     os.Getenv("DATABASE_URI"),
		DatabaseProject: os.Getenv("DATABASE_PROJECT"),
		JWTKey:          os.Getenv("JWT_KEY"),
	}
	return cfg
}
