package controller

import (
	"go-api/internal/db"
	"go-api/internal/usecase"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	usecase *usecase.ProductUseCase 
}

func NewProductController(uc *usecase.ProductUseCase) *ProductController {
	return &ProductController{usecase: uc}
}


// CreateProduct godoc
// @Summary      Cria um novo produto
// @Description  Cria um produto com nome, descrição, preço e data de venda
// @Tags         produtos
// @Accept       json
// @Produce      json
// @Param        produto body db.CreateProductParams true "Dados do produto"
// @Success      201 {object} db.Product
// @Failure      400 {object} map[string]string
// @Router       /product [post]
func (pc *ProductController) CreateProduct(ctx echo.Context) error {
	var req db.CreateProductParams

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	

	product, err := pc.usecase.CreateProduct(ctx.Request().Context(), req) 
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	
	product.DataVenda = req.DataVenda
	
	return ctx.JSON(http.StatusCreated, product)
}


// GetProduct godoc
// @Summary      Lista todos os produtos
// @Description  Retorna todos os produtos cadastrados
// @Tags         produtos
// @Produce      json
// @Success      200 {array} db.Product
// @Failure      500 {object} map[string]string
// @Router       /products [get]
func (pc *ProductController) GetProduct(ctx echo.Context) error {
	product, err := pc.usecase.ListProduct(ctx.Request().Context()) 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, product)
}


// GetProductById godoc
// @Summary      Busca um produto por ID
// @Description  Retorna os dados de um produto específico
// @Tags         produtos
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200 {object} db.Product
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /product/{id} [get]
func (pc *ProductController) GetProductById(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product, err := pc.usecase.GetProduct(ctx.Request().Context(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if product.ID == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
	}

	return ctx.JSON(http.StatusOK, product)
}


// UpdateProduct godoc
// @Summary      Atualiza um produto
// @Description  Atualiza os dados de um produto existente
// @Tags         produtos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Param        produto body db.UpdateProductParams true "Dados atualizados"
// @Success      200 {object} db.Product
// @Failure      400 {object} map[string]string
// @Router       /product/{id} [put]
func (pc *ProductController) UpdateProduct(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var req db.UpdateProductParams
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	} 

	req.ID = int32(id)

	product, err := pc.usecase.UpdateProduct(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, product)
}


// DeleteProduct godoc
// @Summary      Remove um produto
// @Description  Exclui um produto do banco de dados
// @Tags         produtos
// @Param        id path int true "ID do produto"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /product/{id} [delete]
func (pc *ProductController) DeleteProduct(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = pc.usecase.DeleteProduct(ctx.Request().Context(), int32(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}