package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

// GetAllBookings godoc
// @Summary Mostra todas as reservas
// @Description Retorna um JSON com todas as reservas
// @Tags Bookings
// @Produce json
// @Success 200 {array} models.BookingParsed
// @Router /bookings [get]
// @Failure 404 {object} models.ErrorResponse "Nenhuma reserva encontrada"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
func (controller *BookingsController) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, ret := controller.bookingsService.GetAllBookings()

	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
		fmt.Printf("INFO: Bookings found\n")
	case svc.BookingNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "bookings not found"})
		fmt.Printf("INFO: Bookings not found\n")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")
	}

}
