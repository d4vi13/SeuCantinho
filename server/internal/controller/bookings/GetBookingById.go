package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

func (controller *BookingsController) GetBookingById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	booking, ret := controller.bookingsService.GetBookingById(id)

	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
		fmt.Printf("INFO: Booking %d found\n", id)
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
