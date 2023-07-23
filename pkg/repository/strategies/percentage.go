package strategies

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/peer"
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
    var activationFlag string
    if strategy.Execute {
        activationFlag = "active"
    } else {
        activationFlag = "inactive"
    }
    key := strings.Join([]string{strategy.Type, strategy.Tag, strconv.Itoa(strategy.Percentage), activationFlag}, "-")
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

    // Update every peer with the image
    for _, peer := range peers {
        var keyValue models.Crdt
        currentValue, err := crdt.GetCrdtValue(peer)
        // It does not exist
        if err != nil {
            // Register 
            keyValue = models.Crdt{
                Key:   peer,
                Value: strategy.Tag,
            }
        } else {
            // Update
            currentTags := strings.Split(currentValue.Value, ",")
            if tagDoesNotExist(strategy.Tag, currentTags) {
                currentTags = append(currentTags, strategy.Tag)
            }
            updatedTags := strings.Join(currentTags, ",")
            keyValue = models.Crdt{
                Key:   peer,
                Value: updatedTags,
            }
        }

        byteKeyValue, err := json.Marshal(keyValue)
        if err != nil {
            return err
        }
        crdt.PostCrdtKeyValue(byteKeyValue)
    }
    
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
