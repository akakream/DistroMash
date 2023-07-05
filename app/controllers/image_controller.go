package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akakream/DistroMash/app/models"
	"github.com/akakream/DistroMash/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// PostImage uploads a multi-platform docker image to ipfs and get the cid
// @Description Upload a multi-platform docker image to ipfs and get the cid.
// @Summary upload a multi-platform docker image to ipfs and get the cid
// @Tags Image
// @Accept json
// @Produce json
// @Param crdt body models.Image true "Post Image"
// @Success 200 {object} models.Image
// @Router /api/v1/image [post]
func PostImage(c *fiber.Ctx) error {
	cidTagPair, err := uploadImage2IPFS(c.Body())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Add the cid to the CRDT key value store
	crdtPayload, err := json.Marshal(models.Crdt{Key: cidTagPair.Name, Value: cidTagPair.Cid})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	err = postCrdtKeyValue(crdtPayload)
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

func uploadImage2IPFS(imageName []byte) (*models.Image, error) {
	url := fmt.Sprintf("http://%s/image", utils.Multiplatform2ipfsURL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(imageName))
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
		apiErr, err := getErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data models.Image
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
