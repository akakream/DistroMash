package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/akakream/DistroMash/app/models"
	"github.com/akakream/DistroMash/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func getPeersList() (*models.Peers, error) {
	url := fmt.Sprintf("http://%s/peers", utils.Libp2pURL)
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
		apiErr, err := getErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data models.Peers
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func GetPeersListUI(c *fiber.Ctx) error {
	data, err := getPeersList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	peers := strings.Split(data.Peers, ",")

	return c.Render("peers", fiber.Map{
		"Peers": peers,
	}, "base")
}
