package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

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
// @Router /api/v1/strategy/{key} [get]
func GetStrategy(c *fiber.Ctx) error {
	data, err := getCrdtValue(c.Params("key"))
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    strategy, err := parseStrategyFromKey(data.Key)    
    if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
    }

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      false,
		"strategy": strategy,
	})
}

// PostStrategy posts the strategy.
// @Description Post the strategy.
// @Summary post the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param strategy body models.Strategy true "Post Strategy"
// @Success 200 {object} models.Strategy
// @Router /api/v1/strategy [post]
func PostStrategy(c *fiber.Ctx) error {
	var strategy models.Strategy
	if err := json.Unmarshal(c.Body(), &strategy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    if err := checkInput(&strategy); err != nil {
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
		k, v, err := strategies.ProcessStrategyPercentage(&strategy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		key = k
		value = v
	case strategies.StrategySomethingelse:
		fmt.Println("StrategySomethingelse")
	}

	// Call crdt api and register the strategy
	strategyKeyValue := models.Crdt{
		Key:   key,
		Value: value,
	}
    fmt.Printf("Key: %s Value: %s", key, value)

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

// PutStrategy updates the strategy.
// @Description Update the strategy.
// @Summary update the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param strategy body models.Strategy true "Put Strategy"
// @Success 200 {object} models.Strategy
// @Router /api/v1/strategy [put]
func PutStrategy(c *fiber.Ctx) error {
	var strategy models.Strategy
	if err := json.Unmarshal(c.Body(), &strategy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    if err := checkInput(&strategy); err != nil {
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
		k, v, err := strategies.ProcessStrategyPercentage(&strategy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		key = k
		value = v
	case strategies.StrategySomethingelse:
		fmt.Println("StrategySomethingelse")
	}

	// Call crdt api and register the strategy
	strategyKeyValue := models.Crdt{
		Key:   key,
		Value: value,
	}
    fmt.Printf("Key: %s Value: %s", key, value)

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

func getStrategyList(c *fiber.Ctx) ([]models.Strategy, error) {
	data, err := getCrdtList()
	if err != nil {
        return nil, err
	}
    stypes, err := strategies.GetStrategyTypes()
    if err != nil {
        return nil, err
    }

    var existingStrategies []models.Strategy
    for _, pair := range data {
        strategy, err := parseStrategyFromKey(pair.Key)
        if err != nil {
            continue
        }
        if contains(stypes, strategy.Type) {
            existingStrategies = append(existingStrategies, *strategy)
        }
    }

    return existingStrategies, nil
}

func parseStrategyFromKey(key string) (*models.Strategy, error) {
    keyFields := strings.Split(key, "-")
    var strategy models.Strategy
    for i, field := range keyFields {
        if i == 0 {
            strategy.Type = field
        } else if i == 1 {
            strategy.Tag = field
        } else if i == 2 {
            percentage, err := strconv.Atoi(field)
            if err != nil {
                return nil, err
            }
            strategy.Percentage = percentage
        } else if i == 3 {
            if field == "active" {
                strategy.Execute = true
            } else if field == "inactive"{
                strategy.Execute = false
            }
        }
    }
    return &strategy, nil
}

func contains(slice []string, key string) bool {
    for _, s := range slice {
        if s == key {
            return true
        }
    }
    return false
}

func checkInput(strategy *models.Strategy) error {
    if strategy.Percentage > 100 || strategy.Percentage < 0 {
        return errors.New("Percentage must be between 0 and 100")
    }
    return nil
}

// DeleteStrategy deletes the strategy.
// @Description Delete the strategy.
// @Summary delete the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param name path string true "Delete Strategy"
// @Success 200 {object} models.Strategy
// @Router /api/v1/strategy/{key} [delete]
func DeleteStrategy(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      false,
		"strategy": "fake",
	})
}
