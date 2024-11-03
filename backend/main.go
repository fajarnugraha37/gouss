package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/rewrite"
	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/fajarnugraha37/gouss/config"
)

func main() {
	app := fiber.New()

	// static assets
	app.Get("/", static.New("./assets/public/index.html"))
	app.Use([]string{"/favicon.png", "/favicon.ico"}, static.New("./assets/public/favicon.png"))
	app.Use("_app", static.New("./assets/public/_app", static.Config{ Compress: true }))
	
	// api
	apiGroup := app.Group("/api")
	apiGroup.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(
		[]string{
			"/*.html", "/*.htm",
			"/*.ico", "/*.png", "/*.webp", 
			"/*.js", "/*.css", 
			"/*.json", "/*.xml",
		}, 
		static.New("./assets/public", static.Config{ Compress: true }),
	)
	app.Use(
		"/*",
		rewrite.New(rewrite.Config{
			Rules: map[string]string{"/*": "/$1.html"},
			Next: func(c fiber.Ctx) bool {
				if c.Is("json") || c.Route().Method != "GET" {
					return true
				}

				return false
			},
		}),
		static.New("./assets/public", static.Config{Compress: true}),
	)

	address := config.AppProperty.App.Host + ":" + strconv.Itoa(config.AppProperty.App.Port)
	log.Fatal(app.Listen(address))
}
