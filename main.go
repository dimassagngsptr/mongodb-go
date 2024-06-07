package main

import (
	"mongodb-go/src/configs"
	"mongodb-go/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, world",
		})
	})
	configs.Connect()
	routes.Router(app)
	app.Listen(":8080")
}
