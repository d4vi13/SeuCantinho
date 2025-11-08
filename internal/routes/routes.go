package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {

	// B. Obtem todas as reservas
	mux.HandleFunc("GET /booking/{$}", GetAllBookings)

	// B. Obtem uma reserva especifícia
	mux.HandleFunc("GET /booking/{id}", GetBookingById)

	// B. Cancela uma reserva especifíca
	mux.HandleFunc("POST /booking/cancel/{id}", CancelBookingById)

	// B. Obtem todos os espaços (Reservados e Disponíveis)
	mux.HandleFunc("GET /space/{$}", GetAllSpaces)

	// B. Obtem um espaço específico
	mux.HandleFunc("GET /space/{id}", GetSpaceById)

	// B. Reserva um espaço
	mux.HandleFunc("POST /space/book", BookSpace)

	// A. Cria um espaço
	mux.HandleFunc("POST /space/create", CreateSpace)

	// C. Efetua um pagamento
	mux.HandleFunc("POST /pix/{key}", MakeFullPix)

	// C. Efetua o pagamento adiantado do sinal
	mux.HandleFunc("POST /pix/signal/{key}", MakePartialPix)

	mux.HandleFunc("POST /user", CreateUser)
}
