package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

// GetUserBookings godoc
// @Summary Lista de todas as reservas de um usuário
// @Description Retorna um array com todas as reservas do usuário, especificado pelo ID
// @Tags Bookings
// @Produce json
// @Param id path int true "ID usuário"
// @Success 200 {array} models.BookingParsed "Lista de reservas do usuário"
// @Failure 404 {object} models.ErrorResponse "Usuário ou Reserva não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno do Servidor"
// @Router /users/{id}/bookings [get]
func (controller *BookingsController) GetUserBookings(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	bookings, ret := controller.bookingsService.GetUserBookings(id)
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
		fmt.Printf("INFO: User %d bookings were found\n", id)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "user not found"})
		fmt.Printf("INFO: User %d not found\n", id)
	case svc.BookingNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "bookings not found"})
		fmt.Printf("INFO: Bookings %d not found\n", id)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
