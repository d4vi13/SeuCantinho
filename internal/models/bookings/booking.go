package models

type Booking struct {
	Id      int
	UserId  int
	SpaceId int
	PixId   int
	Start   int64
	End     int64
}
