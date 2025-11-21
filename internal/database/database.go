package database

import (
	"database/sql"
	"fmt"
	"os"

  _ "github.com/lib/pq"
)

var pool *sql.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	var err error
	pool, err = sql.Open("postgres", dsn)
	return err
}

func Query(q string, args ...any) (*sql.Rows, error) {
	return pool.Query(q, args...)
}
