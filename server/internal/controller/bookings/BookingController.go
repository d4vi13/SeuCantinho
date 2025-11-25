package bookings

import (
	"github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/services/users"
)

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
}

func (controller *BookingsController) Init() {
	controller.bookingsService.Init()
}
