package main

import (
	"log"
    "strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	
	"github.com/fajarnugraha37/gouss/config"
)

func main() {
	app := fiber.New()

	app.Use("/favicon.png", static.New("./assets/public/favicon.png"))
	app.Use("/_app", static.New("./assets/public/_app", static.ConfigDefault, static.Config{Compress: true}))
	app.Use([]string{"/", "/index.html"}, static.New("./assets/public/index.html"))
	app.Get("/hello-world", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	address := config.AppProperty.App.Host + ":" + strconv.Itoa(config.AppProperty.App.Port)
	log.Fatal(app.Listen(address))
}
