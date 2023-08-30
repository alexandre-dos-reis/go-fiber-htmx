package main

import (
	"go-fiber-htmx/src/engine"
	"go-fiber-htmx/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views:       engine.GetEngine(),
		ViewsLayout: "layouts/main",
	})

	routes.SetupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
