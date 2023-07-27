package controllers

import (
	"github.com/akakream/DistroMash/pkg/repository/ipfs"
	"github.com/gofiber/fiber/v2"
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
