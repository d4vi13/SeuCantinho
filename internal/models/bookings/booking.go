package models

import (
  "time"
  "errors"
)

type Booking struct {
	Id      int
	UserId  int
	SpaceId int
	//PixId   int
	Start   int64 //Unix Time
	End     int64 // Unix Time
}

func (b *Booking) Validate() error {
    if b.Start <= 0 || b.End <= 0 {
        return errors.New("timestamps must be positive")
    }

    if b.Start >= b.End {
        return errors.New("start must be before end")
    }

    now := time.Now().Unix()
    if b.Start <= now {
        return errors.New("start must be in the future")
    }

    return nil
}
