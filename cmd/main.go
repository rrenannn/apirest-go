package main

import (
	"log"
	"go-api/internal/db"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro na conexão com o banco de dados: %v", err)
	}
	defer dbConn.Close()

	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"message": "pong"})
	})

	server.Start(":8080")
}