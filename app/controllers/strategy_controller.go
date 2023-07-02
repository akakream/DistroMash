package controllers

import (
	"github.com/akakream/DistroMash/app/models"
	"github.com/gofiber/fiber/v2"
)

// GetStrategyList gets the whole strategies.
// @Description Get all strategies.
// @Summary get all strategies
// @Tags Strategy
// @Accept json
// @Produce json
// @Success 200 {array} models.Strategy
// @Router /api/v1/strategy [get]
func GetStrategyList(c *fiber.Ctx) error {
	data, err := getStrategyList(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      false,
		"strategies": data,
	})
}

// GetStrategy gets the strategy.
// @Description Get the strategy.
// @Summary get the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param name path string true "Name of Strategy"
// @Success 200 {object} models.Strategy
// @Router /api/v1/strategy/{name} [get]
func GetStrategy(c *fiber.Ctx) error {
	return c.SendString("Strategy A")
}

// PostStrategy posts the strategy.
// @Description Post the strategy.
// @Summary post the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param name body string true "Name of Strategy"
// @Success 200 {object} models.Strategy
// @Router /api/v1/strategy [post]
func PostStrategy(c *fiber.Ctx) error {
	return c.SendString("Strategy A")
}

// GetRegisteredStrategyList gets the registered strategies.
// @Description Get registered strategies.
// @Summary get registered strategies
// @Tags Strategy
// @Accept json
// @Produce json
// @Success 200 {array} models.Strategy
// @Router /api/v1/strategy/registered [get]
func GetRegisteredStrategyList(c *fiber.Ctx) error {
	return c.SendString("Registered Strategy")
}

func getStrategyList(c *fiber.Ctx) ([]models.Strategy, error) {
	return []models.Strategy{{Name: "Strategy1"}, {Name: "Strategy2"}}, nil
}

func GetStrategyListUI(c *fiber.Ctx) error {
	data, err := getStrategyList(c)
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Render("strategy", fiber.Map{
		"Strategy": data,
	}, "base")
}
