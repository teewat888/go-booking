package config

import "os"

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func strEnvOrPanic(key, description string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("envvar is not defined: " + key)
}
