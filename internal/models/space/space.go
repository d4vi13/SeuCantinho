package models

type Photo struct {
	img []byte
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
