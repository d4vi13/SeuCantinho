package users

import "database/sql"

type UsersRepository struct {
	DB *sql.DB
}

func (repository *UsersRepository) Init(db *sql.DB) {
	repository.DB = db
}
