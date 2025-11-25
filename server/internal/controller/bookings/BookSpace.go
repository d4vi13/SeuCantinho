package bookings

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/d4vi13/SeuCantinho/server/internal/services/users"
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
  
  ret := controller.usersService.AuthenticateUser(req.Username, req.Password)
  switch ret {
  case users.UserNotFound:
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"erro": "user not found"})
    fmt.Printf("INFO: User %s not found\n", req.Username)
    return
  case users.WrongPassword:
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]string{"erro": "wrong password"})
    fmt.Printf("INFO: Wrong Password for user %s given\n", req.Username)
    return
  }

  fmt.Printf("INFO: User %s authenticated\n", req.Username)

  userId := controller.usersService.GetUserId(req.Username)
  if userId == -1 {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"erro": "user not found"})
    fmt.Printf("INFO: User %s not found\n", req.Username)
    return
  }

  // TODO check if space id is valid

  // if succeds returns the id of the booking created
  var bookingId int
  bookingId, ret = controller.bookingsService.BookSpace(userId, req.SpaceId, req.Start, req.End)

  /*
  if ret == bookings.BookingCreated {
    // creates a payment for the booking created
    var paymentId int
    paymentId, ret := controller.paymentsService.CreatePayment(id, value)
    if ret == payments.PaymentCreated {

    } else {

    }
  }

} else {
  */
  switch ret {
  case bookings.BookingCreated:
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "New Booking: %d\n", bookingId)
  case bookings.BookingConflict:
    w.WriteHeader(http.StatusConflict)
    fmt.Fprintf(w, "Booking Conflict")
  case bookings.InternalError:
    w.WriteHeader(http.StatusInternalServerError)
  case bookings.BadBooking:
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Invalid Booking")

  }

  // TODO use booking id to create a payment 
}
