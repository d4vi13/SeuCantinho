package payments

import (
	"github.com/d4vi13/SeuCantinho/server/internal/services/payments"
)

type PaymentRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
  Id       int    `json:"id"`
  Value    int32   `json:space`
}

type PaymentsController struct {
	paymentsService payments.PaymentsService
}

func (controller *PaymentsController) Init() {
	controller.paymentsService.Init()
}
