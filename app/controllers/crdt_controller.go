package controllers

import "github.com/gofiber/fiber/v2"

// GetCrdtList gets the whole CRDT store.
// @Description Get all CRDT key-value pairs.
// @Summary get all crdt key-value pairs
// @Tags Crdts
// @Accept json
// @Produce json
// @Success 200 {array} models.Crdt
// @Router /api/v1/crdt [get]
func GetCrdtList(c *fiber.Ctx) error {
	return c.SendString("CRDTS")
}

// GetCrdt gets the CRDT value by key.
// @Description Get the CRDT value by key.
// @Summary get crdt value by given key
// @Tags Crdt
// @Accept json
// @Produce json
// @Param key path string true "Key of Value"
// @Success 200 {array} models.Crdt
// @Router /api/v1/crdt/{key} [get]
func GetCrdtValue(c *fiber.Ctx) error {
	return c.SendString("CRDTS")
}
