package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/akakream/DistroMash/docs"
)

func SwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
