package routes

import (
	"mongodb-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {
	c.Get("/users", controllers.GetUser)
	c.Post("/users", controllers.CreateData)
	c.Put("/user/:id", controllers.UpdateData)
	c.Delete("/user/:id", controllers.DeleteData)
}
