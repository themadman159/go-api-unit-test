package productrepo

import (
	"errors"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	var product model.Product
	err := r.Database.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &product, nil
}
