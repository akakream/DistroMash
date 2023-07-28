package models

type Strategy interface {
	Process() (string, string, error)
}

// Strategy struct to describe a percentage strategy.
type StrategyPayload struct {
	Type       string `db:"type"       json:"type"`
	Tag        string `db:"tag"        json:"tag"`
	Percentage int    `db:"percentage" json:"percentage"`
	Target     string `db:"target"     json:"target"`
	Execute    bool   `db:"execute"    json:"execute"`
}
