package app

import "github.com/gofiber/fiber/v2"

func GetCrdtList(c *fiber.Ctx) error {
	return c.SendString("CRDTS")
}
