package productrepo

import "go-api/pkg/model"

func (r *ProductRepository) GetAllProducts() (*[]model.Product, error) {
	var products []model.Product
	err := r.Database.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}
