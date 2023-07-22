package strategies

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/akakream/DistroMash/models"
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
        executeStrategy(strategy, value)
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

func executeStrategy(strategy *models.Strategy, peerValues string) error {
    // peers := strings.Split(peerValues, ",")
    time.Sleep(time.Second * 5)
    return nil
}
