package models

type Space struct {
	Id         int
	Location   string
	Substation string
	Price      int64
	Capacity   int
	Img        []byte
}
