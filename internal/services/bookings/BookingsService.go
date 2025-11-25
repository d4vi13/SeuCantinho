package bookings

import (
	models "github.com/d4vi13/SeuCantinho/internal/models/bookings"
	"github.com/d4vi13/SeuCantinho/internal/repository/bookings"
	"log"
)

const (
  BookingCreated = iota
  BookingConflict
  BadBooking
  InternalError
  Success
)

type BookingsService struct {
	bookingsRepository bookings.BookingsRepository
}

func (service *BookingsService) Init() {
	service.bookingsRepository.Init()
}

func (service *BookingsService) BookSpace(userId int, spaceId int, start int64, end int64) (int, int) {

  booking := &models.Booking{
    UserId: userId,
		SpaceId: spaceId,
		Start: start,
		End: end,
	}

  err := booking.Validate() 
  if err != nil {
    log.Println(err)
    return -1, BadBooking 
  }

  conflict, err := service.bookingsRepository.CheckBookingConflicts(booking)
  if err != nil {
    log.Println(err)
    return -1, InternalError 
  }

  if conflict {
    log.Println("INFO: Booking Conflict")
    return -1, BookingConflict 
  }

	id, err := service.bookingsRepository.Insert(booking)
	if err != nil {
    log.Println("ERROR: Failed to insert new booking")
		return -1, InternalError 
	}

	return id, Success
}
