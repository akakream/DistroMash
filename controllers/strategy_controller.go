package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/strategies"
)

// GetStrategyList gets the whole strategies.
// @Description Get all strategies.
// @Summary get all strategies
// @Tags Strategy
// @Accept json
// @Produce json
// @Success 200 {array} models.StrategyPayload
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
// @Param key path string true "Key of Strategy"
// @Success 200 {object} models.StrategyPayload
// @Router /api/v1/strategy/{key} [get]
func GetStrategy(c *fiber.Ctx) error {
	data, err := crdt.GetCrdtValue(c.Params("key"))
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	stypes, err := strategies.GetStrategyTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	strategy, err := parseStrategyFromKey(data.Key, stypes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":    false,
		"strategy": strategy,
	})
}

// PostStrategy posts the strategy.
// @Description Post the strategy.
// @Summary post the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param strategy body models.StrategyPayload true "Post Strategy"
// @Success 200 {object} models.StrategyPayload
// @Router /api/v1/strategy [post]
func PostStrategy(c *fiber.Ctx) error {
	var strategyPayload models.StrategyPayload
	if err := json.Unmarshal(c.Body(), &strategyPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	strategyPayload, err := checkInput(strategyPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var strategy models.Strategy

	var key string
	var value string
	// TODO: DO THIS PART MAYBE IN GOROUTINE???
	switch sType := strategyPayload.Type; sType {
	case strategies.StrategyPercentageType:
		var s strategies.StrategyPercentage
		s.Nametag = strategyPayload.Nametag
		s.Type = strategyPayload.Type
		s.Execute = strategyPayload.Execute
		s.Percentage = strategyPayload.Percentage
		strategy = &s
	case strategies.StrategyTargetType:
		var s strategies.StrategyTarget
		s.Nametag = strategyPayload.Nametag
		s.Type = strategyPayload.Type
		s.Execute = strategyPayload.Execute
		s.Target = strategyPayload.Target
		strategy = &s
	}

	key, value, err = strategy.Process()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := registerStrategyToCRDT(key, value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"crdts": "fakedata",
	})
}

func deleteInactiveStrategyIfExists(key string) error {
	keySplit := strings.Split(key, "-")
	activationOrder := (keySplit[len(keySplit)-1] == "active")
	keySplit[len(keySplit)-1] = "inactive"
	if activationOrder {
		crdt.DeleteCrdtKeyValue(strings.Join(keySplit, "-"))
	}
	return nil
}

func registerStrategyToCRDT(key string, value string) error {
	err := deleteInactiveStrategyIfExists(key)
	if err != nil {
		return err
	}
	// Call crdt api and register the strategy
	strategyKeyValue := models.Crdt{
		Key:   key,
		Value: value,
	}
	fmt.Printf("Key: %s Value: %s", key, value)

	byteStrategyKeyValue, err := json.Marshal(strategyKeyValue)
	if err != nil {
		return err
	}
	crdt.PostCrdtKeyValue(byteStrategyKeyValue)
	return nil
}

// PutStrategy updates the strategy.
// @Description Update the strategy.
// @Summary update the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param strategy body models.StrategyPayload true "Put Strategy"
// @Success 200 {object} models.StrategyPayload
// @Router /api/v1/strategy [put]
func PutStrategy(c *fiber.Ctx) error {
	return PostStrategy(c)
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

func getStrategyList(c *fiber.Ctx) ([]models.StrategyPayload, error) {
	data, err := crdt.GetCrdtList()
	if err != nil {
		return nil, err
	}
	stypes, err := strategies.GetStrategyTypes()
	if err != nil {
		return nil, err
	}

	var existingStrategies []models.StrategyPayload
	for _, pair := range data {
		strategy, err := parseStrategyFromKey(pair.Key, stypes)
		if err != nil {
			continue
		}
		existingStrategies = append(existingStrategies, *strategy)
	}

	return existingStrategies, nil
}

func parseKeyFromStrategyPayload(strategy *models.StrategyPayload) (string, error) {
	var target string
	var activity string
	if strategy.Type == "percentage" {
		target = strconv.Itoa(strategy.Percentage)
	} else {
		target = strategy.Target
	}
	if strategy.Execute {
		activity = "active"
	} else {
		activity = "inactive"
	}
	fields := []string{strategy.Type, strategy.Nametag, target, activity}
	key := strings.Join(fields, "-")
	return key, nil
}

func parseStrategyFromKey(key string, stypes []string) (*models.StrategyPayload, error) {
	keyFields := extractTagForEdgeCases(strings.Split(key, "-"))
	var strategy models.StrategyPayload

	// Check if strategy
	if !contains(stypes, keyFields[0]) {
		return nil, errors.New("not a strategy")
	}

	for i, field := range keyFields {
		if i == 0 {
			strategy.Type = field
		} else if i == 1 {
			strategy.Nametag = field
		} else if i == 2 && keyFields[0] == "percentage" {
			percentage, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			strategy.Percentage = percentage
		} else if i == 2 && keyFields[0] == "target" {
			strategy.Target = field
		} else if i == 3 {
			if field == "active" {
				strategy.Execute = true
			} else if field == "inactive" {
				strategy.Execute = false
			}
		}
	}
	return &strategy, nil
}

func extractTagForEdgeCases(keyFields []string) []string {
	if len(keyFields) == 4 {
		return keyFields
	}

	var updatedKeyFields []string
	lenKeyFields := len(keyFields)
	updatedKeyFields[0] = keyFields[0]
	updatedKeyFields[2] = keyFields[lenKeyFields-2]
	updatedKeyFields[3] = keyFields[lenKeyFields-1]
	updatedKeyFields[1] = strings.Join(keyFields[1:(lenKeyFields-1)], "")

	return updatedKeyFields
}

func contains(slice []string, key string) bool {
	for _, s := range slice {
		if s == key {
			return true
		}
	}
	return false
}

func checkInput(strategy models.StrategyPayload) (models.StrategyPayload, error) {
	if strategy.Percentage > 100 || strategy.Percentage < 0 {
		return strategy, errors.New("Percentage must be between 0 and 100")
	}
	if strategy.Type != strategies.StrategyPercentageType &&
		strategy.Type != strategies.StrategyTargetType {
		return strategy, errors.New("Type must be either percentage or target")
	}
	nametag := strings.Split(strategy.Nametag, ":")
	if len(nametag) == 1 {
		strategy.Nametag = strategy.Nametag + ":latest"
	}
	return strategy, nil
}

// DeleteStrategy deletes the strategy.
// @Description Delete the strategy.
// @Summary delete the strategy
// @Tags Strategy
// @Accept json
// @Produce json
// @Param name path string true "Delete Strategy"
// @Success 200 {object} models.StrategyPayload
// @Router /api/v1/strategy/{key} [delete]
func DeleteStrategy(c *fiber.Ctx) error {
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
		"error":    false,
		"strategy": "fake",
	})
}
