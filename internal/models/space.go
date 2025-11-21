package models

import (
	"image"
)

type Photo struct {
	img image.Image
}

type Space struct {
	id         int
	location   string
	substation string
	price      float64
	capacity   int
	isAdmin    bool
	photo      Photo
}
