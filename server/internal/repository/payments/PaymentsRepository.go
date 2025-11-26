package payments

import (
	"database/sql"
	"errors"

	"github.com/d4vi13/SeuCantinho/server/internal/database"
	models "github.com/d4vi13/SeuCantinho/server/internal/models/payments"
)

type PaymentsRepository struct{}

func (repository *PaymentsRepository) Init() {}

func (repository *PaymentsRepository) Insert(payment *models.Payment) (int, error) {
	query := `INSERT INTO payments (id, totalValue, payedValue) VALUES ($1, $2, $3) RETURNING id;`

	var id int
	err := database.QueryRow(query, payment.Id, payment.TotalValue, payment.PayedValue).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repository *PaymentsRepository) GetPaymentById(id int) (*models.Payment, error) {
	query := `SELECT id, totalValue, payedValue FROM payments WHERE id = $1`
	payment := &models.Payment{}

	row := database.QueryRow(query, id)

	err := row.Scan(&payment.Id, &payment.TotalValue, &payment.PayedValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return payment, nil
}
