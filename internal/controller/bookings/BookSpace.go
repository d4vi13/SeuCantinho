package bookings

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (controller *BookingsController) BookSpace(w http.ResponseWriter, r *http.Request) {
	var req BookSpaceRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	booking, ret := controller.bookingsService.BookSpace(req.Username, req.Space, req.Start, req.End)

	switch ret {
	case bookings.BookingCreated:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "New Booking: %+v\n", booking)
	case bookings.UnavailableDate:
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User already exists")
	case booking.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
