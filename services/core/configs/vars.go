package configs

import "os"

type Config struct {
	HTTPServerConfig *HTTPServerConfig
}

type HTTPServerConfig struct {
	Endpoint string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{
		HTTPServerConfig: &HTTPServerConfig{
			Endpoint: getEnv("HTTP_ENDPOINT", ":3001"),
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
