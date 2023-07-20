package routes

import (
	"github.com/akakream/DistroMash/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	v1.Get("/strategy", controllers.GetStrategyList)
	v1.Get("/strategy/:name", controllers.GetStrategy)
	v1.Post("/strategy", controllers.PostStrategy)

	v1.Get("/dig/:tag", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("tag"))
	})

	v1.Get("/crdt", controllers.GetCrdtList)
	v1.Get("/crdt/:key", controllers.GetCrdtValue)
	v1.Post("/crdt", controllers.PostCrdtValue)
	v1.Delete("/crdt/:key", controllers.DeleteCrdtValue)

	v1.Post("/image", controllers.PostImage)
}
