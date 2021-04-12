package configs

import (
	"os"
)

const (
	ModeDev        = "dev"
	ModeProduction = "prod"
)

type MailConfig struct {
	Sender    string
	Password  string
	Host      string
	Port      string
	Receivers []string
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}
type EnvConfig struct {
	ServerMode      string
	Port            string
	DatabaseUri     string
	DatabaseProject string
	JWTKey          string
	Mail            MailConfig
	AWS             AWSConfig
}

func Load() EnvConfig {
	cfg := EnvConfig{
		ServerMode: os.Getenv("SERVER_MODE"),
		Port:       os.Getenv("APP_PORT"),
		Mail: MailConfig{
			Sender:   os.Getenv("MAIL_SENDER"),
			Password: os.Getenv("MAIL_PASSWORD"),
			Host:     os.Getenv("MAIL_HOST"),
			Port:     os.Getenv("MAIL_PORT"),
			Receivers: []string{
				os.Getenv("MAIL_RECEIVER"),
			},
		},
		DatabaseUri:     os.Getenv("DATABASE_URI"),
		DatabaseProject: os.Getenv("DATABASE_PROJECT"),
		JWTKey:          os.Getenv("JWT_KEY"),
		AWS: AWSConfig{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:          os.Getenv("AWS_REGION"),
		},
	}

	return cfg
}
