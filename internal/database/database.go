package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	// Normaliza porta: remove prefixo ':' caso o usuário tenha colocado ':5432'
	port = strings.TrimPrefix(port, ":")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica conexão imediatamente; fecha em caso de erro para não vazar conexões
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
