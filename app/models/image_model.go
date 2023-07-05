package models

// Image struct to describe a docker image name.
type Image struct {
	Name string `db:"name" json:"name"`
}

// Image struct to describe a docker image name-cid pair.
type ImageWithCID struct {
	Name string `db:"name" json:"name"`
	Cid  string `db:"cid" json:"cid"`
}
