package usecase

import (
	"context"
	"errors"
	"go-api/internal/db"
	"go-api/internal/repository"
	"strconv"
	"strings"
	"time"
)

type ProductUseCase struct {
	repo repository.ProductRepositoryInterface
}

func NewProductUseCase(repo repository.ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func validateCreateProduct(arg db.CreateProductParams) error {
	if strings.TrimSpace(arg.Nome) == "" {
        return errors.New("nome do produto é obrigatório")
    }

    if len(strings.TrimSpace(arg.Descricao)) < 10 {
        return errors.New("descrição deve ter pelo menos 10 caracteres")
    }

    precoFloat, err := strconv.ParseFloat(arg.Preco, 64)
    if err != nil {
        return errors.New("preço inválido")
    }

    if precoFloat < 0 {
        return errors.New("preço não pode ser negativo")
    }

    if arg.DataVenda.After(time.Now()) {
        return errors.New("data de venda não pode ser futura")
    }

    return nil
}

func validateUpdateProduct(arg db.UpdateProductParams) error {
	if strings.TrimSpace(arg.Nome) == "" {
        return errors.New("nome do produto é obrigatório")
    }

    if len(strings.TrimSpace(arg.Descricao)) < 10 {
        return errors.New("descrição deve ter pelo menos 10 caracteres")
    }

    precoFloat, err := strconv.ParseFloat(arg.Preco, 64)
    if err != nil {
        return errors.New("preço inválido")
    }

    if precoFloat < 0 {
        return errors.New("preço não pode ser negativo")
    }

    if arg.DataVenda.After(time.Now()) {
        return errors.New("data de venda não pode ser futura")
    }

    return nil
}


func (pu *ProductUseCase) ListProduct(ctx context.Context) ([]db.Product, error) {
	return pu.repo.ListProduct(ctx)
}

func (pu *ProductUseCase) GetProduct(ctx context.Context, id int32) (db.Product, error) {
	return pu.repo.GetProduct(ctx, id)
}

func (pu *ProductUseCase) CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error) {
	err := validateCreateProduct(arg)
	if err != nil {
		return db.Product{}, err
	}

	return pu.repo.CreateProduct(ctx, arg)
}

func (pu *ProductUseCase) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error) {
	err := validateUpdateProduct(arg)
	if err != nil {
		return db.Product{}, err
	}

	return pu.repo.UpdateProduct(ctx, arg)
}

func (pu *ProductUseCase) DeleteProduct(ctx context.Context, id int32) error {
	return pu.repo.DeleteProduct(ctx, id)
}