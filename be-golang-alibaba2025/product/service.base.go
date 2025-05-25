package product

import "context"

type Service interface {
	GetAll(ctx context.Context) (products []*Product, err error)
	GetByID(ctx context.Context, id int) (product *Product, err error)
	AddProduct(ctx context.Context, req AddProductRequest) (product *Product, err error)
	UpdateProduct(ctx context.Context, id int, req UpdateProductRequest) (err error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) *service {
	return &service{repo: repository}
}
