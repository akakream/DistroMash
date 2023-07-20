package routes

import (
	"github.com/akakream/DistroMash/controllers"
	"github.com/gofiber/fiber/v2"
)

func UiRoutes(app *fiber.App) {
	app.Get("/", controllers.GetHomeUI)

	app.Get("/crdt", controllers.GetCrdtListUI)
	app.Get("/strategy", controllers.GetStrategyListUI)
	app.Get("/peers", controllers.GetPeersListUI)
}
