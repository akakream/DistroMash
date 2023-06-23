package models

// Crdt struct to describe a crdt key-value pair.
type Crdt struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}
