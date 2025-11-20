package users

import (
	"database/sql"
	"errors"
	"log"

	"github.com/d4vi13/SeuCantinho/internal/database"
	models "github.com/d4vi13/SeuCantinho/internal/models/users"
)

type UsersRepository struct {
	DB *sql.DB
}

func (repository *UsersRepository) Init() {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal("DB error:", err)
	}
	repository.DB = conn
}

func (repository *UsersRepository) Insert(user *models.User) (int, error) {
	query := `
		INSERT INTO users (username, pass_hash, is_admin)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var id int
	err := repository.DB.QueryRow(query, user.Username, user.PassHash, user.IsAdmin).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *UsersRepository) GetUserByName(username string) (*models.User, error) {
	query := `SELECT id, username, pass_hash, is_admin FROM users WHERE username = $1`
	user := &models.User{}

	row := repository.DB.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.PassHash, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
