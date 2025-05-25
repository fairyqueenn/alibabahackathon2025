package product

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(ctx context.Context) (products []*Product, err error)
	FindByID(ctx context.Context, id int) (product *Product, err error)
	UpsertProduct(ctx context.Context, product *Product) (err error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{DB: db}
}
