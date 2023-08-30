package main

import (
	"go-fiber-htmx/src/engine"
	"go-fiber-htmx/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	engine := engine.GetEngine()

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	routes.SetupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
