package bookings

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/d4vi13/SeuCantinho/server/internal/services/users"
  "github.com/d4vi13/SeuCantinho/server/internal/services/bookings"
  "github.com/d4vi13/SeuCantinho/server/internal/services/payments"
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

  fmt.Printf("INFO: Booking %d created \n", bookingId)

  if ret == bookings.Success {
    var paymentId int
    var value int64
    value, _ = controller.spaceService.ComputeBookingPrice(req.SpaceId, req.End - req.Start)
    fmt.Printf("INFO: Booking Value %f\n", value)
    paymentId, ret = controller.paymentsService.CreatePayment(bookingId, value)
    switch ret {
    case payments.Success:
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode(map[string]int{"id": paymentId})
      fmt.Printf("INFO: Booking %d created succesfuly\n", bookingId)
    default:
      controller.bookingsService.CancelBookingById(bookingId)
      w.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
      fmt.Printf("ERROR: Unknown Status\n")
    }
  } else {
    fmt.Printf("INFO: other if\n")
    switch ret {
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
    }
  }
}
