package routes

import (
	"net/http"

	"github.com/d4vi13/SeuCantinho/server/internal/controller/space"
	"github.com/d4vi13/SeuCantinho/server/internal/controller/users"
	"github.com/d4vi13/SeuCantinho/server/internal/controller/bookings"
)

func RegisterRoutes(mux *http.ServeMux) {

	var usersController users.UsersController
	var spaceController space.SpaceController
	var bookingsController bookings.BookingsController
	usersController.Init()
	spaceController.Init()
	bookingsController.Init()

	// // B. Obtem todas as reservas
	mux.HandleFunc("GET /bookings", bookingsController.GetAllBookings)

	// // B. Obtem uma reserva especifícia
	mux.HandleFunc("GET /bookings/{id}", bookingsController.GetBookingById)

	// // B. Cancela uma reserva especifíca
	mux.HandleFunc("DELETE /bookings/{id}", bookingsController.CancelBookingById)

	// // B. Reserva um espaço
	 mux.HandleFunc("POST /bookings", bookingsController.BookSpace)

	mux.HandleFunc("GET /space/{$}", spaceController.GetAllSpaces)

	mux.HandleFunc("GET /space/{id}", spaceController.GetSpaceById)

	mux.HandleFunc("POST /space", spaceController.CreateSpace)

	mux.HandleFunc("PUT /space/{id}", spaceController.UpdateSpace)

	mux.HandleFunc("DELETE /space/{id}", spaceController.DeleteSpace)

	// // C. Efetua um pagamento
	// mux.HandleFunc("POST /pix/{key}", MakeFullPix)

	// // C. Efetua o pagamento adiantado do sinal
	// mux.HandleFunc("POST /pix/signal/{key}", MakePartialPix)

	mux.HandleFunc("POST /users", usersController.CreateUser)

	mux.HandleFunc("POST /login", usersController.UserLogin)

}
