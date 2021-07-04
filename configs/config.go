package configs

import (
	"os"
)

const (
	ModeDev        = "dev"
	ModeProduction = "prod"
)

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
	ACL             string
}
type EnvConfig struct {
	ServerMode      string
	Port            string
	DatabaseUri     string
	DatabaseProject string
	JWTKey          string
	AWS             AWSConfig
}

func Load() EnvConfig {
	cfg := EnvConfig{
		ServerMode:      os.Getenv("SERVER_MODE"),
		Port:            os.Getenv("APP_PORT"),
		DatabaseUri:     os.Getenv("DATABASE_URI"),
		DatabaseProject: os.Getenv("DATABASE_PROJECT"),
		JWTKey:          os.Getenv("JWT_KEY"),
		AWS: AWSConfig{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:          os.Getenv("AWS_REGION"),
			Bucket:          os.Getenv("AWS_BUCKET_NAME"),
			ACL:             "public-read",
		},
	}

	return cfg
}
