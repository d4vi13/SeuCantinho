package bookings

import (
	"github.com/d4vi13/SeuCantinho/internal/services/bookings"
)

type BookSpaceRequest struct {
	Username string `json:"username"`
	Space    uint32 `json:space`
	Start    int64  `json:start`
	End      int64  `json:end`
}

type BookingsController struct {
	BookingsService bookings.BookingsService
}

func (controller *BookingsController) Init() {
	controller.bookings.Service.Init()
}
