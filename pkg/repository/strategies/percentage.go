package strategies

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/ipfs"
	"github.com/akakream/DistroMash/pkg/repository/peer"
	cid "github.com/ipfs/go-cid"
)

func ProcessStrategyPercentage(strategy *models.Strategy) (string, string, error) {
	// Calculate involving peers
	key, err := constructKey(strategy)
	if err != nil {
		return "", "", err
	}
	value, err := constructValue(float64(strategy.Percentage))
	if err != nil {
		return "", "", err
	}
    
    // Execute strategy if execute field is true
    if strategy.Execute {
        executeStrategy(strategy, key, value)
    }

	return key, value, nil
}

func constructKey(strategy *models.Strategy) (string, error) {
    tag := strategy.Tag
    // Reverse Lookup if CID instead of tag
    _, err := cid.Decode(strategy.Tag)
    if err == nil {
        crdtEntry, err := crdt.GetCrdtValue(strategy.Tag)
        if err != nil {
            log.Println("THE CID IS UNKNOWN!!!")
        }
        tag = crdtEntry.Value
    }

    var activationFlag string
    if strategy.Execute {
        activationFlag = "active"
    } else {
        activationFlag = "inactive"
    }
    key := strings.Join([]string{strategy.Type, tag, strconv.Itoa(strategy.Percentage), activationFlag}, "-")
	return key, nil
}

func constructValue(percentage float64) (string, error) {
	peers, err := peer.GetPeersList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return "", err
	}

	selectedPeers, err := randomNPercentOfPeers(peers.Peers, percentage, 1)
	if err != nil {
		return "", err
	}

    var selectedPeerIDs []string
    for _, p := range selectedPeers {
        selectedPeerIDs = append(selectedPeerIDs, p.ID)
    }

	return strings.Join(selectedPeerIDs, ","), nil
}

func randomNPercentOfPeers(peers []models.Peer, percentage float64, seed int64) ([]models.Peer, error) {
	rand.Seed(seed)

	selectedCount := int(float64(len(peers)) * percentage / 100)
	selected := make([]models.Peer, selectedCount)

	for i := 0; i < selectedCount; i++ {
		index := rand.Intn(len(peers))
		selected[i] = peers[index]
		peers = append(peers[:index], peers[index+1:]...)
	}

	return selected, nil
}

func executeStrategy(strategy *models.Strategy, key string, peerValues string) error {
    peers :=strings.Split(peerValues, ",")
    tagToCidChan := make(chan string)
    go resolveTagToCID(strategy.Tag, tagToCidChan)
    go updatePeers(peers, tagToCidChan)
    
    return nil
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
