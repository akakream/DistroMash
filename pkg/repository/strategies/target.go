package strategies

import (
	"errors"
	"fmt"
	"log"
	"strings"

	cid "github.com/ipfs/go-cid"

	"github.com/akakream/DistroMash/pkg/repository/crdt"
	"github.com/akakream/DistroMash/pkg/repository/peer"
)

// Strategy struct to describe a target strategy.
type StrategyTarget struct {
	Type    string `db:"type"    json:"type"`
	Tag     string `db:"tag"     json:"tag"`
	Target  string `db:"target"  json:"target"`
	Execute bool   `db:"execute" json:"execute"`
}

func (strategy *StrategyTarget) Process() (string, string, error) {
	// Calculate involving peers
	key, err := strategy.constructKey()
	if err != nil {
		return "", "", err
	}
	value, err := strategy.constructValue()
	if err != nil {
		return "", "", err
	}

	// Execute strategy if execute field is true
	if strategy.Execute {
		strategy.executeStrategy(key, value)
	}

	return key, value, nil
}

func (strategy *StrategyTarget) constructKey() (string, error) {
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
	key := strings.Join(
		[]string{strategy.Type, tag, strategy.Target, activationFlag},
		"-",
	)
	return key, nil
}

func (strategy *StrategyTarget) constructValue() (string, error) {
	err := targetInPeerList(strategy)
	if err != nil {
		return "", err
	}
	return strategy.Target, nil
}

func (strategy *StrategyTarget) executeStrategy(
	key string,
	peerValues string,
) error {
	peers := strings.Split(peerValues, ",")
	tagToCidChan := make(chan string)
	go resolveTagToCID(strategy.Tag, tagToCidChan)
	go updatePeers(peers, tagToCidChan)

	return nil
}

func targetInPeerList(strategy *StrategyTarget) error {
	peers, err := peer.GetPeersList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return fmt.Errorf("failed to get peers list: %w", err)
	}

	// Check if the peer exists in the peer list
	for _, p := range peers.Peers {
		if strategy.Target == p.ID {
			return nil
		}
	}
	return errors.New("peer is not in the peer list")
}
