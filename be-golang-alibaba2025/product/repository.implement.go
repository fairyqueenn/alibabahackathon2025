package product

import "context"

func (r *repository) UpsertProduct(ctx context.Context, product *Product) (err error) {
	err = r.DB.Save(product).Error
	return
}

func (r *repository) FindByID(ctx context.Context, id int) (product *Product, err error) {
	err = r.DB.Where("id = ?", id).First(&product).Error
	return
}

func (r *repository) FindAll(ctx context.Context) (products []*Product, err error) {
	err = r.DB.Find(&products).Error
	return
}
