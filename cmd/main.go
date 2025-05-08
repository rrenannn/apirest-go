package main

import (
	"fmt"
	"go-api/internal/controller"
	"go-api/internal/db"
	"go-api/internal/repository"
	"go-api/internal/usecase"
	"log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-api/docs"
	"github.com/labstack/echo/v4"
)

// @title           API REST Gerenciamento de produtos
// @version         1.0
// @description     CRUD para gerenciamento de produtos.
// @host            localhost:8080

func main() {
	server := echo.New()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro na conexão com o banco de dados: %v", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	ProductRepository := repository.NewProductRepository(queries)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)


	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.GET("/products", ProductController.GetProduct)
	server.GET("/product/:id", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product/:id", ProductController.UpdateProduct)
	server.DELETE("/product/:id", ProductController.DeleteProduct)


	fmt.Println("Rodando...")
	server.Start(":8080")
}