package models

// Strategy struct to describe a strategy.
type Strategy struct {
	Type string `db:"type" json:"type"`
	Tag string `db:"tag" json:"tag"`
	Percentage int `db:"percentage" json:"percentage"`
    Execute bool `db:"execute" json:"execute"`
}
