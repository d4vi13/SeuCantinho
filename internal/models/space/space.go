package models

type Space struct {
	Id         int
	Location   string
	Substation string
	Price      float64
	Capacity   int
	Img        []byte
}
