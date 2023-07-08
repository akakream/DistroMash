package models

// Peer struct to describe a peer.
type Peer struct {
	Addr string `db:"addr" json:"addr"`
}

type Peers struct {
	Peers string `json:"peers"`
}
