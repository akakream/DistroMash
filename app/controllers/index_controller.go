package controllers

import "github.com/gofiber/fiber/v2"

func GetHomeUI(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello world",
	}, "base")
}
