package routes

import (
	"github.com/akakream/DistroMash/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func IndexRoutes(app *fiber.App) {
	app.Get("/", controllers.GetHome)
}
