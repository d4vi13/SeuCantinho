package bookings

import (
	models "github.com/d4vi13/SeuCantinho/internal/models/bookings"
	"github.com/d4vi13/SeuCantinho/internal/repository/bookings"
)

const (
  BookingCreated = iota
  BookingConflict
  InternalError
  Success
)

type BookingsService struct {
	bookingsRepository bookings.BookingsRepository
}

func (service *BookingsService) Init() {
	service.bookings Repository.Init()
}

func (service *bookingsService) BookSpace(userId int, spaceId int, start int64, end int64) (int, int) {

	booking = &models.Bookings{
    UserId: userId,
		SpaceId: spaceId,
		Start: start,
		End: end
	}

  err := service.BookingsRepository.CheckBookingConflicts(booking)
  if err != nil {
    return -1, BookingConflict 
  }

	id, err := service.BookingsRepository.Insert(user)
	if err != nil {
		return -1, InternalError 
	}

	return id, Success
}
