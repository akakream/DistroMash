package models

// Image struct to describe a docker image name-cid pair.
type Image struct {
	Name string `db:"name" json:"name"`
	Cid  string `db:"cid" json:"cid"`
}
