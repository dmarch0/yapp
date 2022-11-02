package server

import (
	"yapp/core/configs"
	"yapp/core/internal/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func StartServer() {
	config := configs.GetConfig()
	app := fiber.New()
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.HTTPServerConfig.CookieSecret,
	}))
	routes.AddEndpoints(app)
	app.Listen(config.HTTPServerConfig.Endpoint)
}
