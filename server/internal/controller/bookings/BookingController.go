package bookings

import (
	"github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/services/users"
	"github.com/d4vi13/SeuCantinho/server/internal/services/space"
	"github.com/d4vi13/SeuCantinho/server/internal/services/payments"
)

type  AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BookSpaceRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
  SpaceId    int    `json:space`
	Start    int64  `json:start` // Unix Time
	End      int64  `json:end` // Unix Time
}

type BookingsController struct {
	bookingsService bookings.BookingsService
	usersService users.UsersService
	paymentsService payments.PaymentsService
	spaceService space.SpaceService
}

func (controller *BookingsController) Init() {
	controller.bookingsService.Init()
}
