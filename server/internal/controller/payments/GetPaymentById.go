package payments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/payments"
)

// GetPaymentById godoc
// @Summary Retorna a entidade de um pagamento específico
// @Description Retorna um JSON com os dados do pagamento referente ao ID informado
// @Tags Payments
// @Produce json
// @Param id path int true "ID do pagamento"
// @Success 200 {object} models.Payment "Dados do pagamento encontrados"
// @Failure 404 {object} map[string]string "Pagamento não encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /payments/{id} [get]
func (controller *PaymentsController) GetPaymentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	payment, ret := controller.paymentsService.GetPaymentById(id)

	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Success:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payment)
		fmt.Printf("INFO: Payment %d found\n", id)
	case svc.NotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "payment not found"})
		fmt.Printf("INFO: payment %d not found\n", id)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")
	}

}
