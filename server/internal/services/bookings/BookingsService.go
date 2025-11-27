package bookings

import (
	models "github.com/d4vi13/SeuCantinho/server/internal/models/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/services/users"
	"github.com/d4vi13/SeuCantinho/server/internal/services/payments"
	"github.com/d4vi13/SeuCantinho/server/internal/services/space"
	"github.com/d4vi13/SeuCantinho/server/internal/repository/bookings"
	"log"
)

const (
  BookingCreated = iota
  BookingConflict
  BookingNotFound
  BadBooking
  SpaceNotFound
  PaymentCreationFailed
  InternalError
  UserNotFound
  WrongPassword
  Unauthorized
  Success
)

type BookingsService struct {
	bookingsRepository bookings.BookingsRepository
	paymentsService payments.PaymentsService
	spaceService space.SpaceService
	usersService users.UsersService
}

func (service *BookingsService) Init() {
	service.bookingsRepository.Init()
}

func (service *BookingsService) BookSpace(username string, password string, spaceId int, start int64, end int64) (int, int) {

  ret := service.usersService.AuthenticateUser(username, password)
	if ret == users.UserNotFound {
		log.Printf("BookingsService: User Not Found\n")
		return -1, UserNotFound
	}

	if ret == users.WrongPassword {
		log.Printf("BookingsService: Wrong Password\n")
		return -1, WrongPassword
	}

  userId := service.usersService.GetUserId(username)
	if userId == -1 {
		log.Printf("BookingsService: User Not Found\n")
		return -1, UserNotFound
	}

  var value int64
  value, ret = service.spaceService.ComputeBookingPrice(spaceId, end - start)
  if ret == space.SpaceNotFound {
		log.Printf("BookingsService: Space Not Found\n")
		return -1, SpaceNotFound
	}

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

  _, ret = service.paymentsService.CreatePayment(id, value)
  if ret != payments.Success {
	  service.bookingsRepository.Delete(id)
    return -1, PaymentCreationFailed 
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

func (service *BookingsService) IsBookingOwner(userId, bookingId int) bool {

	booking, err := service.bookingsRepository.GetBookingById(bookingId)
	if err != nil {
		log.Printf("%+v\n", err)
		return false 
	}

  if booking.UserId != userId {
    return false
  }

  return true
}

func (service *BookingsService) CancelBookingById(username string, password string, bookingId int) int {

  ret := service.usersService.AuthenticateUser(username, password)
	if ret == users.UserNotFound {
		log.Printf("BookingsService: User Not Found\n")
		return UserNotFound
	}

	if ret == users.WrongPassword {
		log.Printf("BookingsService: Wrong Password\n")
		return WrongPassword
	}

  userId := service.usersService.GetUserId(username)
	if userId == -1 {
		log.Printf("BookingsService: User Not Found\n")
		return UserNotFound
	}

  if !service.usersService.UserIsAdmin(username) && !service.IsBookingOwner(userId, bookingId) {
    return Unauthorized
  }

	err := service.bookingsRepository.Delete(bookingId)
	if err != nil {
		log.Printf("%+v\n", err)
		return BookingNotFound
	}

	return Success
}

func (service *BookingsService) GetUserBookings(userId int, username string, password string)  ([]models.Booking, int) {

  _, ret := service.usersService.GetUserById(userId)
  if ret == users.UserNotFound {
		log.Printf("BookingsService: User Not Found\n")
		return nil, UserNotFound
	}

  ret = service.usersService.AuthenticateUser(username, password)
	if ret == users.UserNotFound {
		log.Printf("BookingsService: User Not Found\n")
		return nil, UserNotFound
	}

	if ret == users.WrongPassword {
		log.Printf("BookingsService: Wrong Password\n")
		return nil, WrongPassword
	}

  requesterId := service.usersService.GetUserId(username)
	if userId == -1 {
		log.Printf("BookingsService: User Not Found\n")
		return nil, UserNotFound
	}

  if !service.usersService.UserIsAdmin(username) && !(userId != requesterId) {
    return nil, Unauthorized
  }

  bookings, err := service.bookingsRepository.GetUserBookings(userId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, BookingNotFound
	}

	return bookings, Success
}

