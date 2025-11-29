package payments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/payments"
)

func (controller *PaymentsController) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest

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

	ret := controller.paymentsService.MakePayment(id, req.Value)

	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"info": "payment made"})
		fmt.Printf("INFO: payment %d was payed\n", id)
	case svc.InvalidPayment:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid payment attempted"})
		fmt.Printf("ERROR: invalid payment %d\n", id)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
