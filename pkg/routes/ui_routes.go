package routes

import (
	"github.com/akakream/DistroMash/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func UiRoutes(app *fiber.App) {
	ui := app.Group("/ui")

	ui.Get("/crdt", controllers.GetCrdtListUI)
	ui.Get("/strategy", controllers.GetStrategyListUI)
}
