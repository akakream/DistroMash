package controllers

import "github.com/gofiber/fiber/v2"

// GetStrategyList gets the whole strategies.
// @Description Get all strategies.
// @Summary get all strategies
// @Tags Strategy
// @Accept json
// @Produce json
// @Success 200 {array} models.Strategy
// @Router /strategy [get]
func GetStrategyList(c *fiber.Ctx) error {
	return c.SendString("Strategy")
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
// @Router /strategy [post]
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
// @Router /strategy/registered [get]
func GetRegisteredStrategyList(c *fiber.Ctx) error {
	return c.SendString("Registered Strategy")
}
