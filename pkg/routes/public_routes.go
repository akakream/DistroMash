package routes

import "github.com/gofiber/fiber/v2"

func PublicRoutes(app *fiber.App) {
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	v1.Get("/strategy/:name?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	v1.Post("/strategy/:name?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	v1.Get("/dig/:tag?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("tag"))
	})

	v1.Get("/crdt", Get)
}