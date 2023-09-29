package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/controllers"
)

func UiRoutes(app *fiber.App) {
	app.Get("/", controllers.GetStrategyListUI)

	app.Get("/crdt", controllers.GetCrdtListUI)
	app.Get("/strategy", controllers.GetStrategyListUI)
	app.Get("/peers", controllers.GetPeersListUI)
}
