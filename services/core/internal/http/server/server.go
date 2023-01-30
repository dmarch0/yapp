package server

import (
	"yapp/core/configs"
	"yapp/core/internal/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func StartServer() {
	config := configs.GetConfig()
	app := fiber.New()
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.HTTPServerConfig.CookieSecret,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.HTTPServerConfig.AllowedOrigin,
	}))
	routes.AddEndpoints(app)
	app.Listen(config.HTTPServerConfig.Endpoint)
}
