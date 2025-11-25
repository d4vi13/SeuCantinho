package bookings

import (
	models "github.com/d4vi13/SeuCantinho/server/internal/models/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/repository/bookings"
	"log"
)

const (
  BookingCreated = iota
  BookingConflict
  BookingNotFound
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

func (service *BookingsService) GetAllBookings() ([]models.Booking, int) {
	bookings, err := service.bookingsRepository.GetAllBookings()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, InternalError
	}

	if len(bookings) == 0 {
		log.Printf("Bookings not found\n")
		return nil, BookingNotFound
	}

	return bookings, Success
}


func (service *BookingsService) GetBookingById(bookingId int) (*models.Booking, int) {

	booking, err := service.bookingsRepository.GetBookingById(bookingId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, BookingNotFound
	}

	return booking, Success
}


