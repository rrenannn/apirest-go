package main

import (
	"fmt"
	"go-api/internal/controller"
	"go-api/internal/db"
	"go-api/internal/repository"
	"go-api/internal/usecase"
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

	queries := db.New(dbConn)

	ProductRepository := repository.NewProductRepository(queries)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)
	


	fmt.Println("Rodando...")
	server.Start(":8080")
}