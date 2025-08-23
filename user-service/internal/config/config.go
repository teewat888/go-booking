package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type (
	HTTPConfig struct {
		Port string
		Host string
	}

	Config struct {
		ServiceId  string
		Env        string
		HTTP       HTTPConfig
		DbUrl      string
		NatsUrl    string
		AuthSecret string
	}
)

func FromEnv() Config {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Cannot load the .env file - err: %v", err))
	}
	return Config{
		ServiceId: getEnv("SERVICE_ID", "msgo-boilerplate"),
		DbUrl:     strEnvOrPanic("DB_URL", "DB url connection string"),
		Env:       getEnv("ENV", "dev"),
		HTTP: HTTPConfig{
			Port: getEnv("HTTP_PORT", "3000"),
			Host: getEnv("HTTP_HOST", "0.0.0.0"),
		},
		NatsUrl:    strEnvOrPanic("NATS_URL", "NATS url  string"),
		AuthSecret: strEnvOrPanic("AUTH_SECRET", "Auth secret string"),
	}
}
