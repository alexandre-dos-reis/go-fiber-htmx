package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/home", fiber.Map{})
	})
	app.Get("/contact/1/edit", func(c *fiber.Ctx) error {
		return c.Render("pages/contact", fiber.Map{})
	})
}
