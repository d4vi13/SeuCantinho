package bookings

import (
	"database/sql"
	"errors"

	"github.com/d4vi13/SeuCantinho/server/internal/database"
	models "github.com/d4vi13/SeuCantinho/server/internal/models/bookings"
)

type BookingsRepository struct{}

func (repository *BookingsRepository) Init() {}

func (repository *BookingsRepository) CheckBookingConflicts(booking *models.Booking) (bool, error) {
  var exists int
	query := `SELECT 1 FROM bookings WHERE spaceId = $1 AND bookingStart < $3 AND bookingEnd > $2 LIMIT 1;`

  err := database.QueryRow(query,booking.SpaceId, booking.Start, booking.End).Scan(&exists)
  if err != nil {
    if err == sql.ErrNoRows {
      // no conflict
      return false, nil
    }

    //internal error
    return true, errors.New("internal db error")
  }

  // conflict
  return true, nil 
}

func (repository *BookingsRepository) Insert(booking *models.Booking) (int, error) {
	// Statement para inserir um novo espaço
	query := `INSERT INTO bookings (spaceId, userId, bookingStart, bookingEnd) VALUES ($1, $2, $3, $4) RETURNING id;`

	// Preenche os dados e obtém o ID
	var id int
	err := database.QueryRow(query, booking.SpaceId, booking.UserId, booking.Start, booking.End).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repository *BookingsRepository) GetAllBookings() ([]models.Booking, error) {
	bookings := make([]models.Booking, 0)

	query := `SELECT id, spaceId, userId, bookingStart, bookingEnd FROM bookings ORDER BY id`
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(&booking.Id, &booking.SpaceId, &booking.UserId, &booking.Start, &booking.End)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (repository *BookingsRepository) GetBookingById(id int) (*models.Booking, error) {
	query := `SELECT id, spaceId, userId, bookingStart, bookingEnd FROM bookings WHERE id = $1`
	booking := &models.Booking{}

	row := database.QueryRow(query, id)

	err := row.Scan(&booking.Id, &booking.SpaceId, &booking.UserId, &booking.Start, &booking.End)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}

	return booking, nil
}

