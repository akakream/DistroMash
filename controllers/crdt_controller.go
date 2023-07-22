package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/pkg/repository/crdt"
)

// GetCrdtList gets the whole CRDT store.
// @Description Get all CRDT key-value pairs.
// @Summary get all crdt key-value pairs
// @Tags Crdt
// @Accept json
// @Produce json
// @Success 200 {array} models.Crdt
// @Router /api/v1/crdt [get]
func GetCrdtList(c *fiber.Ctx) error {
	data, err := crdt.GetCrdtList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"crdts": data,
	})
}

// GetCrdt gets the CRDT value by key.
// @Description Get the CRDT value by key.
// @Summary get crdt value by given key
// @Tags Crdt
// @Accept json
// @Produce json
// @Param key path string true "Key of Value"
// @Success 200 {object} models.Crdt
// @Router /api/v1/crdt/{key} [get]
func GetCrdtValue(c *fiber.Ctx) error {
	data, err := crdt.GetCrdtValue(c.Params("key"))
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"crdt":  data,
	})
}

// PostCrdt posts a CRDT key-value pair.
// @Description Post a CRDT key-value pair.
// @Summary post crdt key-value pair
// @Tags Crdt
// @Accept json
// @Produce json
// @Param crdt body models.Crdt true "Post Crdt"
// @Success 200 {object} models.Crdt
// @Router /api/v1/crdt [post]
func PostCrdtValue(c *fiber.Ctx) error {
	err := crdt.PostCrdtKeyValue(c.Body())
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "success",
	})
}

func GetCrdtListUI(c *fiber.Ctx) error {
	data, err := crdt.GetCrdtList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Render("crdt", fiber.Map{
		"Crdt": data,
	}, "base")
}

// DeleteCrdt deletes a CRDT key-value pair.
// @Description Delete a CRDT key-value pair.
// @Summary delete crdt key-value pair
// @Tags Crdt
// @Accept json
// @Produce json
// @Param key path string true "Key of Value"
// @Success 200 {object} models.Crdt
// @Router /api/v1/crdt/{key} [delete]
func DeleteCrdtValue(c *fiber.Ctx) error {
	err := crdt.DeleteCrdtKeyValue(c.Params("key"))
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "success",
	})
}
