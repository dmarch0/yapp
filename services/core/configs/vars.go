package configs

import "os"

type Config struct {
	HTTPServerConfig *HTTPServerConfig
	MongoDBConfig    *MongoDBConfig
}

type HTTPServerConfig struct {
	Endpoint     string
	CookieSecret string
}

type MongoDBConfig struct {
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
			CookieSecret: getEnv("COOKIE_SECRET", "ThisIsNotSecret"),
		},
		MongoDBConfig: &MongoDBConfig{
			Uri: getEnv("MONGO_URI", "mongodb://mongo:27017"),
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
