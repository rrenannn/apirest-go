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

func (pc *ProductController) CreateProduct(ctx echo.Context) error {
	var req db.CreateProductParams

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	req.DataVenda.Format("2006-01-02")


	product, err := pc.usecase.CreateProduct(ctx.Request().Context(), req) 
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	
	product.DataVenda = req.DataVenda
	
	return ctx.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProduct(ctx echo.Context) error {
	product, err := pc.usecase.ListProduct(ctx.Request().Context()) 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, product)
}

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