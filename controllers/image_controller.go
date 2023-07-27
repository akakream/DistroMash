package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/ipfs"
	"github.com/gofiber/fiber/v2"
)

type jobResult struct {
    Data string
    Error error
}

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
    postResultChan := make(chan jobResult)
    go logPostResult(postResultChan)
    go asyncPostImage(postResultChan, c.Body())

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Upload order is given and being processed.",
	})
}

func logPostResult(postResultChan <-chan jobResult) {
    result := <-postResultChan
    if result.Error != nil {
        log.Println(result.Error)
    } else {
        log.Println(result.Data)
    }
}

func asyncPostImage(postResultChan chan<- jobResult, imageTag []byte) {
    var result jobResult

	cidTagPair, err := ipfs.UploadImage2IPFS(imageTag)
	if err != nil {
        result.Error = err
	}
    
    // Add the cid to the CRDT key value store
    crdtPayload, err := json.Marshal(models.Crdt{Key: cidTagPair.Name, Value: cidTagPair.Cid})
    if err != nil {
        result.Error = err
    }
    err = crdt.PostCrdtKeyValue(crdtPayload)
    if err != nil {
        result.Error = err
    }

    if result.Error == nil {
        result.Data = fmt.Sprintf("Image Name: %s CID: %s", cidTagPair.Name, cidTagPair.Cid)
    }

    postResultChan <- result
}
