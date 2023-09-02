package models

// Image struct to describe a docker image name:tag.
type Image struct {
	Name string `db:"name" json:"name"`
	Tag  string `db:"tag"  json:"tag"`
}

// Image struct to describe a docker image name:tag-cid pair.
type ImageWithCID struct {
	Name string `db:"name" json:"name"`
	Tag  string `db:"tag"  json:"tag"`
	Cid  string `db:"cid"  json:"cid"`
}
