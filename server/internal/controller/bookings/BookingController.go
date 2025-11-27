package bookings

import (
	"github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BookSpaceRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SpaceId  int    `json:"space"`
	Start    int64  `json:"start"`
	End      int64  `json:"end"`
}

type BookingsController struct {
	bookingsService bookings.BookingsService
}

func (controller *BookingsController) Init() {
	controller.bookingsService.Init()
}
