package routes

import "github.com/gofiber/fiber/v2"

func AddHelloEndpoints(router fiber.Router) {
	router.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
