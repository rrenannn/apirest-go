package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	conn := "postgresql://postgres:postgres@localhost:5432/apirest?sslmode=disable"
	return sql.Open("postgres", conn)
}