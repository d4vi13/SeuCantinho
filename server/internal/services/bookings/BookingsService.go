package bookings

import (
	"log"
	"time"

	models "github.com/d4vi13/SeuCantinho/server/internal/models/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/repository/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/services/payments"
	"github.com/d4vi13/SeuCantinho/server/internal/services/space"
	"github.com/d4vi13/SeuCantinho/server/internal/services/users"
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
	paymentsService    payments.PaymentsService
	spaceService       space.SpaceService
	usersService       users.UsersService
}

func (service *BookingsService) Init() {
	service.bookingsRepository.Init()
}

func (service *BookingsService) BookSpace(username string, password string, spaceId int, startDate string, days int) (int, int) {

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

	startParse, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Printf("BookingsService: Error in date parsing\n")
		return -1, InternalError
	}
	endParse := startParse.AddDate(0, 0, days)

	start := startParse.Unix()
	end := endParse.Unix()

	var value int64
	value, ret = service.spaceService.ComputeBookingPrice(spaceId, end-start)
	if ret == space.SpaceNotFound {
		log.Printf("BookingsService: Space Not Found\n")
		return -1, SpaceNotFound
	}

	booking := &models.Booking{
		UserId:  userId,
		SpaceId: spaceId,
		Start:   start,
		End:     end,
	}

	err = booking.Validate()
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

func (service *BookingsService) GetAllBookings() ([]models.BookingParsed, int) {
	bookings := make([]models.BookingParsed, 0)

	repoBookings, err := service.bookingsRepository.GetAllBookings()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, InternalError
	}

	if len(repoBookings) == 0 {
		log.Printf("Bookings not found\n")
		return nil, BookingNotFound
	}

	for _, b := range repoBookings {
		var booking = models.BookingParsed{}

		booking.Id = b.Id
		booking.UserId = b.UserId
		booking.SpaceId = b.SpaceId

		startParsed := time.Unix(b.Start, 0)
		endParsed := time.Unix(b.End, 0)

		booking.StartDate = startParsed.Format("2006-01-02")
		booking.EndDate = endParsed.Format("2006-01-02")

		diff := endParsed.Sub(startParsed)
		booking.Days = int(diff.Hours() / 24)

		bookings = append(bookings, booking)
	}

	return bookings, Success
}

func (service *BookingsService) GetBookingById(bookingId int) (*models.BookingParsed, int) {

	b, err := service.bookingsRepository.GetBookingById(bookingId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, BookingNotFound
	}

	var booking = &models.BookingParsed{}

	booking.Id = b.Id
	booking.UserId = b.UserId
	booking.SpaceId = b.SpaceId

	startParsed := time.Unix(b.Start, 0)
	endParsed := time.Unix(b.End, 0)

	booking.StartDate = startParsed.Format("2006-01-02")
	booking.EndDate = endParsed.Format("2006-01-02")

	diff := endParsed.Sub(startParsed)
	booking.Days = int(diff.Hours() / 24)

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

func (service *BookingsService) GetUserBookings(userId int) ([]models.BookingParsed, int) {
	bookings := make([]models.BookingParsed, 0)

	_, ret := service.usersService.GetUserById(userId)
	if ret == users.UserNotFound {
		log.Printf("BookingsService: User Not Found\n")
		return nil, UserNotFound
	}

	repoBookings, err := service.bookingsRepository.GetUserBookings(userId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, InternalError
	}

	if len(repoBookings) == 0 {
		log.Printf("Bookings not found\n")
		return nil, BookingNotFound
	}

	for _, b := range repoBookings {
		var booking = models.BookingParsed{}

		booking.Id = b.Id
		booking.UserId = b.UserId
		booking.SpaceId = b.SpaceId

		startParsed := time.Unix(b.Start, 0)
		endParsed := time.Unix(b.End, 0)

		booking.StartDate = startParsed.Format("2006-01-02")
		booking.EndDate = endParsed.Format("2006-01-02")

		diff := endParsed.Sub(startParsed)
		booking.Days = int(diff.Hours() / 24)

		bookings = append(bookings, booking)
	}

	return bookings, Success
}
