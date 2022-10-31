package server

import (
	"yapp/core/configs"
	"yapp/core/internal/http/routes"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	config := configs.GetConfig()
	app := fiber.New()
	routes.AddEndpoints(app)
	app.Listen(config.HTTPServerConfig.Endpoint)
}
