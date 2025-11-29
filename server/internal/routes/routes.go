package routes

import (
	"net/http"

	"github.com/d4vi13/SeuCantinho/server/internal/controller/bookings"
	"github.com/d4vi13/SeuCantinho/server/internal/controller/payments"
	"github.com/d4vi13/SeuCantinho/server/internal/controller/space"
	"github.com/d4vi13/SeuCantinho/server/internal/controller/users"
)

func RegisterRoutes(mux *http.ServeMux) {

	var usersController users.UsersController
	var spaceController space.SpaceController
	var bookingsController bookings.BookingsController
	var paymentsController payments.PaymentsController
	usersController.Init()
	spaceController.Init()
	bookingsController.Init()
	paymentsController.Init()

	//Aciona o Controlador de Reservas e mostra todas as reservas
	mux.HandleFunc("GET /bookings/{$}", bookingsController.GetAllBookings)

	//Aciona o Controlador de Reservas e mostra a reserva selecionada pelo ID
	mux.HandleFunc("GET /bookings/{id}", bookingsController.GetBookingById)

	//Aciona o Controlador de Reservas e cancela a reserva selecionada pelo ID
	mux.HandleFunc("DELETE /bookings/{id}", bookingsController.CancelBookingById)

	//Aciona o Controlador de Reservas e adiciona uma nova reserva
	mux.HandleFunc("POST /bookings", bookingsController.BookSpace)

	//Aciona o Controlador de Reservas e mostra todas as reservas de um usuário específico
	mux.HandleFunc("GET /users/{id}/bookings", bookingsController.GetUserBookings)

	//Aciona o Controlador de Pagamentos e executa o pagamento do ID especificado
	mux.HandleFunc("POST /payments/{id}", paymentsController.MakePayment)

	//Aciona o Controlador de Pagamentos e devolve um JSON com a entidade do pagamento
	mux.HandleFunc("GET /payments/{id}", paymentsController.GetPaymentById)

	//Aciona o Controlador de Espaços e devolve um JSON com todos os espaços livres e reservados, mas não informa seus status.
	mux.HandleFunc("GET /space/{$}", spaceController.GetAllSpaces)

	//Aciona o Controlador de Usuários e devolve todos os usuários
	mux.HandleFunc("GET /users/{$}", usersController.GetAllUsers)

	//Aciona o Controlador de Espaços e retorna o espaço pelo ID
	mux.HandleFunc("GET /space/{id}", spaceController.GetSpaceById)

	//Aciona o Controlador de Espaços e adiciona um novo espaço
	mux.HandleFunc("POST /space", spaceController.CreateSpace)

	//Aciona o Controlador de Espaços atualiza o espaço especificado pelo ID
	mux.HandleFunc("PUT /space/{id}", spaceController.UpdateSpace)

	//Aciona o Controlador de Espaços e deleta o espaço especificado pelo ID
	mux.HandleFunc("DELETE /space/{id}", spaceController.DeleteSpace)

	//Aciona o Controlador de Usuários e adiciona um novo usuário
	mux.HandleFunc("POST /users", usersController.CreateUser)

	//Aciona o Controlador de Usuários e realiza o login
	mux.HandleFunc("POST /login", usersController.UserLogin)

	//Aciona o Controlador de Usuários e retorna o usuário especificado pelo ID
	mux.HandleFunc("GET /users/{id}", usersController.GetUserById)

}
