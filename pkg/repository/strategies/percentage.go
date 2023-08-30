package strategies

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"strings"

	cid "github.com/ipfs/go-cid"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/peer"
)

// Strategy struct to describe a percentage strategy.
type StrategyPercentage struct {
	Type       string `db:"type"       json:"type"`
	Tag        string `db:"tag"        json:"tag"`
	Percentage int    `db:"percentage" json:"percentage"`
	Execute    bool   `db:"execute"    json:"execute"`
}

func (strategy *StrategyPercentage) Process() (string, string, error) {
	// Calculate involving peers
	key, err := strategy.constructKey()
	if err != nil {
		return "", "", err
	}
	value, err := strategy.constructValue(float64(strategy.Percentage))
	if err != nil {
		return "", "", err
	}

	// Execute strategy if execute field is true
	if strategy.Execute {
		strategy.executeStrategy(key, value)
	}

	return key, value, nil
}

func (strategy *StrategyPercentage) constructKey() (string, error) {
	tag := strategy.Tag
	// Reverse Lookup if CID instead of tag
	_, err := cid.Decode(strategy.Tag)
	if err == nil {
		crdtEntry, err := crdt.GetCrdtValue(strategy.Tag)
		if err != nil {
			log.Println("THE CID IS UNKNOWN!!!")
			return "", errors.New(
				"The following CID is unknown. Please provide a known CID or a Docker image tag.",
			)
		}
		tag = crdtEntry.Value
	}

	var activationFlag string
	if strategy.Execute {
		activationFlag = "active"
	} else {
		activationFlag = "inactive"
	}
	key := strings.Join(
		[]string{strategy.Type, tag, strconv.Itoa(strategy.Percentage), activationFlag},
		"-",
	)
	return key, nil
}

func (strategy *StrategyPercentage) constructValue(percentage float64) (string, error) {
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

func (strategy *StrategyPercentage) executeStrategy(key string, peerValues string) error {
	peers := strings.Split(peerValues, ",")
	tagToCidChan := make(chan string)
	go resolveTagToCID(strategy.Tag, tagToCidChan)
	go updatePeers(peers, tagToCidChan)

	return nil
}

func randomNPercentOfPeers(
	peers []models.Peer,
	percentage float64,
	seed int64,
) ([]models.Peer, error) {
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
