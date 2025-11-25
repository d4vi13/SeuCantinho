package routes

import (
	"net/http"

	"github.com/d4vi13/SeuCantinho/internal/controller/space"
	"github.com/d4vi13/SeuCantinho/internal/controller/users"
)

func RegisterRoutes(mux *http.ServeMux) {

	var usersController users.UsersController
	var spaceController space.SpaceController
	usersController.Init()
	spaceController.Init()

	// // B. Obtem todas as reservas
	// mux.HandleFunc("GET /booking/{$}", GetAllBookings)

	// // B. Obtem uma reserva especifícia
	// mux.HandleFunc("GET /booking/{id}", GetBookingById)

	// // B. Cancela uma reserva especifíca
	// mux.HandleFunc("POST /booking/cancel/{id}", CancelBookingById)

	// // B. Obtem todos os espaços (Reservados e Disponíveis)
	// mux.HandleFunc("GET /space/{$}", GetAllSpaces)

	mux.HandleFunc("GET /space/{id}", spaceController.GetSpaceById)

	// // B. Reserva um espaço
	// mux.HandleFunc("POST /space/book", BookSpace)

	mux.HandleFunc("POST /space", spaceController.CreateSpace)

	mux.HandleFunc("DELETE /space/{id}", spaceController.DeleteSpace)

	// // C. Efetua um pagamento
	// mux.HandleFunc("POST /pix/{key}", MakeFullPix)

	// // C. Efetua o pagamento adiantado do sinal
	// mux.HandleFunc("POST /pix/signal/{key}", MakePartialPix)

	mux.HandleFunc("POST /users", usersController.CreateUser)

	mux.HandleFunc("POST /login", usersController.UserLogin)

}
