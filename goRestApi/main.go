package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Static("/", "public")
	app.Use(cors.New())

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Este es un usario valido",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}
