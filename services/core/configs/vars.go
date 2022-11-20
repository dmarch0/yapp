package configs

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

type Config struct {
	HTTPServerConfig *HTTPServerConfig
	PostgreSQLConfig *PostgreSQLConfig
}

type HTTPServerConfig struct {
	Endpoint     string
	CookieSecret string
}

type PostgreSQLConfig struct {
	Uri string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{
		HTTPServerConfig: &HTTPServerConfig{
			Endpoint:     getEnv("HTTP_ENDPOINT", ":3001"),
			CookieSecret: getEnv("COOKIE_SECRET", encryptcookie.GenerateKey()),
		},
		PostgreSQLConfig: &PostgreSQLConfig{
			Uri: getEnv("MONGO_URI", "postgresql://admin:admin@postgres:5432/yapp"),
		},
	}
	return config
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
