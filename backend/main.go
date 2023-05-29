package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: `{"ip":${ip}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Static("/", "./public")

	app.Listen(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
