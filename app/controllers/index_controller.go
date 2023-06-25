package controllers

import "github.com/gofiber/fiber/v2"

// GetHome gets the home page.
// @Description Get the home page.
// @Summary get the home page
// @Tags Home
// @Accept json
// @Produce json
// @Success 200 string ok
// @Router / [get]
func GetHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello world",
	}, "base")
}
