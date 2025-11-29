package payments

import (
	"github.com/d4vi13/SeuCantinho/server/internal/services/payments"
)

type PaymentRequest struct {
	Value int64 `json:"value"`
}

type PaymentsController struct {
	paymentsService payments.PaymentsService
}

func (controller *PaymentsController) Init() {
	controller.paymentsService.Init()
}
