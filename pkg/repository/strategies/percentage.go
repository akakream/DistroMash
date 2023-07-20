package strategies

import (
	"math/rand"
	"strings"

	"github.com/akakream/DistroMash/models"
    "github.com/akakream/DistroMash/pkg/repository/peer"
)

func ProcessStrategyPercentage(strategy *models.Strategy) (string, string, error) {
	// Calculate involving peers
	key, err := constructKey()
	if err != nil {
		return "", "", err
	}
	value, err := constructValue(float64(strategy.Percentage))
	if err != nil {
		return "", "", err
	}
	return key, value, nil
}

func constructKey() (string, error) {
	return "", nil
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

	return strings.Join(selectedPeers, ","), nil
}

func randomNPercentOfPeers(peers []string, percentage float64, seed int64) ([]string, error) {
	rand.Seed(seed)

	selectedCount := int(float64(len(peers)) * percentage / 100)
	selected := make([]string, selectedCount)

	for i := 0; i < selectedCount; i++ {
		index := rand.Intn(len(peers))
		selected[i] = peers[index]
		peers = append(peers[:index], peers[index+1:]...)
	}

	return selected, nil
}
