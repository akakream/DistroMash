package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/utils"
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
	data, err := getCrdtList()
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

func getCrdtList() ([]models.Crdt, error) {
	url := fmt.Sprintf("http://%s/crdt", utils.Libp2pURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data []models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
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
	data, err := getCrdtValue(c.Params("key"))
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

func getCrdtValue(key string) (*models.Crdt, error) {
	url := fmt.Sprintf("http://%s/crdt/%s", utils.Libp2pURL, key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
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
	err := postCrdtKeyValue(c.Body())
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

func postCrdtKeyValue(keyValuePair []byte) error {
	url := fmt.Sprintf("http://%s/crdt", utils.Libp2pURL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(keyValuePair))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	return nil
}


func GetCrdtListUI(c *fiber.Ctx) error {
	data, err := getCrdtList()
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
	err := deleteCrdtKeyValue(c.Params("key"))
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

func deleteCrdtKeyValue(key string) error {
	url := fmt.Sprintf("http://%s/crdt/%s", utils.Libp2pURL, key)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	return nil
}
