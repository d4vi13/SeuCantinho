package payments

import (
	models "github.com/d4vi13/SeuCantinho/server/internal/models/payments"
	"github.com/d4vi13/SeuCantinho/server/internal/repository/payments"
	"log"
)

const (
  Success = iota
  NotFound
  InternalError
  InvalidPayment
)

type PaymentsService struct {
	paymentsRepository payments.PaymentsRepository
}

func (service *PaymentsService) Init() {
	service.paymentsRepository.Init()
}

func (service *PaymentsService) CreatePayment(id int, value int64) (int, int) {

  payment := &models.Payment{
    Id: id,
		TotalValue: value,
    PayedValue: 0,
	}

	id, err := service.paymentsRepository.Insert(payment)
	if err != nil {
    log.Println("ERROR: Failed to insert new payment")
		return -1, InternalError 
	}

  log.Printf("INFO: Created Payment with id %d\n", id)
	return id, Success
}

func (service *PaymentsService) GetPaymentById(paymentId int) (*models.Payment, int) {

	payment, err := service.paymentsRepository.GetPaymentById(paymentId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, NotFound
	}

	return payment, Success
}
func (service *PaymentsService) MakePayment(paymentId int, value int64)  int {

  if value < 0 {
    return InvalidPayment
  }

	err := service.paymentsRepository.MakePayment(paymentId, value)
	if err != nil {
		log.Printf("%+v\n", err)
		return InvalidPayment
	}

	return Success
}
