package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	conn := "postgresql://postgres:Guinho%4001@localhost:5432/products_apirest?sslmode=disable"
	return sql.Open("postgres", conn)
}