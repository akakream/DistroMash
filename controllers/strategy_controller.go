package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/strategies"
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
	var strategy *models.Strategy
	if err := json.Unmarshal(c.Body(), strategy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var key string
	var value string
	// TODO: DO THIS PART MAYBE IN GOROUTINE???
	switch sType := strategy.Type; sType {
	case strategies.StrategyPercentage:
		k, v, err := strategies.ProcessStrategyPercentage(strategy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		key = k
		value = v
	case strategies.StrategySomethingelse:
		fmt.Println("WUT")
	}

	// Call crdt api and register the strategy
	strategyKeyValue := models.Crdt{
		Key:   key,
		Value: value,
	}

	byteStrategyKeyValue, err := json.Marshal(strategyKeyValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	postCrdtKeyValue(byteStrategyKeyValue)

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"crdts": "fakedata",
	})
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
	return []models.Strategy{{Type: "Strategy1"}, {Type: "Strategy2"}}, nil
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
