package main

import (
	"fmt"
	"go-api/internal/db"
	"log"

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

	fmt.Println("Rodando...")
	server.Start(":8080")
}