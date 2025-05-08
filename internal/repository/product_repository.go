package repository

import (
	"context"
	"go-api/internal/db"
)

type ProductRepositoryInterface interface{
	ListProduct(ctx context.Context) ([]db.Product, error)
	GetProduct(ctx context.Context, id int32) (db.Product, error)
	CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error)
	UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error)
	DeleteProduct(ctx context.Context, id int32) error
}

var _ ProductRepositoryInterface = (*ProductRepository)(nil)

type ProductRepository struct {
	db *db.Queries
}

func NewProductRepository(db *db.Queries) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) ListProduct(ctx context.Context) ([]db.Product, error) {
	return pr.db.ListProduct(ctx)
}

func (pr *ProductRepository) GetProduct(ctx context.Context, id int32) (db.Product, error) {
	return pr.db.GetProduct(ctx, id)
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error) {
	return pr.db.CreateProduct(ctx, arg)
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error) {
	return pr.db.UpdateProduct(ctx, arg)
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, id int32) error {
	return pr.db.DeleteProduct(ctx, id)
}

