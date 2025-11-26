package users

import (
	"database/sql"
	"errors"

	"github.com/d4vi13/SeuCantinho/server/internal/database"
	models "github.com/d4vi13/SeuCantinho/server/internal/models/users"
)

type UsersRepository struct{}

func (repository *UsersRepository) Init() {}

func (repository *UsersRepository) Insert(user *models.User) (int, error) {
	query := `
		INSERT INTO users (username, pass_hash, is_admin)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	// Insere  o usuário no banco de dados e retorna seu ID
	var id int
	err := database.QueryRow(query, user.Username, user.PassHash, user.IsAdmin).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *UsersRepository) GetUserById(id int) (*models.User, error) {
	query := `SELECT id, username, is_admin FROM users WHERE id = $1`
	user := &models.User{}

	// Busca no banco pelo usuário com id específico
	row := database.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Username, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.PassHash = "hashed_password"
	return user, nil
}

func (repository *UsersRepository) GetUserByName(username string) (*models.User, error) {
	query := `SELECT id, username, pass_hash, is_admin FROM users WHERE username = $1`
	user := &models.User{}

	// Busca no banco pelo usuário com username específico
	row := database.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.PassHash, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
