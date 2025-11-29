package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

func (controller *BookingsController) BookSpace(w http.ResponseWriter, r *http.Request) {
	var req BookSpaceRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	bookingId, ret := controller.bookingsService.BookSpace(req.Username, req.Password, req.SpaceId, req.StartDate, req.Days)
	switch ret {
	case bookings.Success:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": bookingId})
		fmt.Printf("INFO: Booking %d created succesfuly\n", bookingId)
	case bookings.BookingConflict:
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "booking conflict"})
		fmt.Printf("INFO: Booking Conflitc\n")
	case bookings.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		fmt.Printf("ERROR: Internal Server Error\n")
	case bookings.BadBooking:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid booking attempted"})
		fmt.Printf("ERROR: Invalid Booking Attempted\n")
	case bookings.UserNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "user not found"})
		fmt.Printf("INFO: User %s not found\n", req.Username)
	case bookings.SpaceNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "space not found"})
		fmt.Printf("INFO: Space %d not found\n", req.SpaceId)
	case bookings.PaymentCreationFailed:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "payment creation failed"})
		fmt.Printf("INFO: Payment creation failed \n")
	case bookings.WrongPassword:
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"erro": "wrong password"})
		fmt.Printf("INFO: Wrong Password for user %s given\n", req.Username)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")
	}
}
