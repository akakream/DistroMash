package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/pkg/repository/ipfs"
)

// PostImage uploads a multi-platform docker image to ipfs and get the cid
// @Description Upload a multi-platform docker image to ipfs and get the cid.
// @Summary upload a multi-platform docker image to ipfs and get the cid
// @Tags Image
// @Accept json
// @Produce json
// @Param crdt body models.Image true "Post Image"
// @Success 200 {object} models.ImageWithCID
// @Router /api/v1/image [post]
func PostImage(c *fiber.Ctx) error {
	postResultChan := make(chan ipfs.JobResult)
	go ipfs.LogPostResult(postResultChan)
	go ipfs.AsyncPostImage(postResultChan, c.Body())

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Upload order is given and being processed.",
	})
}

// PinCid pins a cid
// @Description Pin cid.
// @Summary pin cid
// @Tags Cid
// @Accept json
// @Produce json
// @Param cid path string true "Cid"
// @Success 200 {object} models.Crdt
// @Router /api/v1/pin/{cid} [post]
func PinCid(c *fiber.Ctx) error {
	err := ipfs.PinCid(c.Params("cid"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "OK",
	})
}
