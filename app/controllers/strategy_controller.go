package controllers

import "github.com/gofiber/fiber/v2"

func GetStrategy(c *fiber.Ctx) error {
	return c.SendString("CRDTS")
}
