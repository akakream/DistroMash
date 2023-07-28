package strategies

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ipfs/go-cid"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/ipfs"
)

const (
	StrategyPercentageType = "percentage"
	StrategyTargetType     = "target"
)

func GetStrategyTypes() ([]string, error) {
	strategies := []string{StrategyPercentageType, StrategyTargetType}
	return strategies, nil
}

func tagDoesNotExist(tag string, currentTags []string) bool {
	for _, currentTag := range currentTags {
		if tag == currentTag {
			return false
		}
	}
	return true
}

func resolveTagToCID(tag string, tagToCidChan chan<- string) {
	// Check if the tag is already CID
	c := tag
	_, err := cid.Decode(tag)
	if err != nil {
		log.Println("Not a CID!")
		val, err := getIfPresentOrPostCRDTstore(tag)
		if err != nil {
			log.Println(err)
		} else {
			c = val
		}
	}

	tagToCidChan <- c
}

func getIfPresentOrPostCRDTstore(tag string) (string, error) {
	// Check if an entry for the tag in the CRDT Store exists
	crdtEntry, err := crdt.GetCrdtValue(tag)
	if err != nil {
		log.Println("The tag-cid mapping does not exist")
		// Download the docker image and push to crdt
		payload := models.Image{
			Name: tag,
		}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return "", err
		}
		postResultChan := make(chan ipfs.JobResult)
		go ipfs.AsyncPostImage(postResultChan, payloadBytes)
		result := <-postResultChan
		if result.Error != nil {
			return "", result.Error
		}
		return result.Data.Cid, nil
	}

	return crdtEntry.Value, nil
}

func updatePeers(peers []string, tagToCidChan <-chan string) {
	cidTag := <-tagToCidChan

	// Update every peer with the image
	for _, peer := range peers {
		var keyValue models.Crdt
		currentValue, err := crdt.GetCrdtValue(peer)
		// It does not exist
		if err != nil {
			// Register
			keyValue = models.Crdt{
				Key:   peer,
				Value: cidTag,
			}
		} else {
			// Update
			currentTags := strings.Split(currentValue.Value, ",")
			if tagDoesNotExist(cidTag, currentTags) {
				currentTags = append(currentTags, cidTag)
				updatedTags := strings.Join(currentTags, ",")
				keyValue = models.Crdt{
					Key:   peer,
					Value: updatedTags,
				}
			} else {
				// If key exists, do not update
				continue
			}
		}

		byteKeyValue, err := json.Marshal(keyValue)
		crdt.PostCrdtKeyValue(byteKeyValue)
	}
}
