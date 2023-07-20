package models

// Peer struct to describe a peer.
type Peer struct {
    ID string `db:"id" json:"id"`
	Addrs []string `db:"addrs" json:"addrs"`
}

type Peers struct {
	Peers []string `json:"peers"`
}
