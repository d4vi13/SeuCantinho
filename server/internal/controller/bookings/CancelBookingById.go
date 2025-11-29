package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

// CancelBookingById godoc
// @Summary Deleta a reserva especifcada pelo ID
// @Description Cancela uma reserva
// @Tags Bookings
// @Produce json
// @Param id path int true "ID da reserva"
// @Param credentials body AuthRequest true "Credenciais do usuário"
// @Success 200 {object} map[string]string "Reserva cancelada"
// @Failure 400 {object} map[string]string "JSON inválido"
// @Failure 401 {object} map[string]string "Senha incorreta ou usuário sem permissão"
// @Failure 404 {object} map[string]string "Usuário ou reserva não encontrada"
// @Failure 500 {object} map[string]string "Erro interno do Servidor"
// @Router /bookings/{id} [delete]
func (controller *BookingsController) CancelBookingById(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest

	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ret := controller.bookingsService.CancelBookingById(req.Username, req.Password, id)
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"info": "deleted"})
		fmt.Printf("INFO: Booking %d was canceled\n", id)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		fmt.Printf("INFO: User %s not found\n", req.Username)
	case svc.WrongPassword:
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong password"})
		fmt.Printf("INFO: Wrong Password for user %s given\n", req.Username)
	case svc.Unauthorized:
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "not owner or admin"})
		fmt.Printf("INFO: User %s doesnt have authority to cancel Booking %d\n", req.Username, id)
	case svc.BookingNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "booking not found"})
		fmt.Printf("INFO: Booking %d not found\n", id)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
