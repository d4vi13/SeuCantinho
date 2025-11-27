package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
)

func (controller *BookingsController) GetUserBookings(w http.ResponseWriter, r *http.Request) {
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

  bookings, ret := controller.bookingsService.GetUserBookings(id, req.Username, req.Password)
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
		fmt.Printf("INFO: User %d bookings were found\n", id)
  case svc.UserNotFound:
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"erro": "user not found"})
    fmt.Printf("INFO: User %s not found\n", req.Username)
  case svc.WrongPassword:
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]string{"erro": "wrong password"})
    fmt.Printf("INFO: Wrong Password for user %s given\n", req.Username)
  case svc.Unauthorized:
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]string{"error": "not owner or admin"})
    fmt.Printf("INFO: User %s doesnt have authority to get user %d bookings\n", req.Username, id)
  case svc.BookingNotFound:
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"error": "booking not found"})
    fmt.Printf("INFO: Booking not found\n", req.Username, id)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
