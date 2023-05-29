package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Static("/", "./public")

	app.Listen(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
