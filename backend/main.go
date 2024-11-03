package main

import (
	"log"

	
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	app.Get("/hello-world", func(c fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
	app.Use([]string{"/", "/index.html"}, static.New("./assets/public/index.html"))
	app.Use("/favicon.png", static.New("./assets/public/favicon.png"))
	app.Use("/_app", static.New("./assets/public/_app", static.ConfigDefault, static.Config{ Compress: true }))

    log.Fatal(app.Listen(":3000"))
}