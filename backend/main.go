package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	app.Use("/favicon.png", static.New("./assets/public/favicon.png"))
	app.Use("/_app", static.New("./assets/public/_app", static.ConfigDefault, static.Config{Compress: true}))
	app.Use([]string{"/", "/index.html"}, static.New("./assets/public/index.html"))
	app.Get("/hello-world", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
