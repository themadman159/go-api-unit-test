package productsvc

import (
	"go-api/pkg/model"

	"gorm.io/gorm"
)

func (s *ProductService) GetByID(id int) (*model.Product, error) {

	product, err := s.ProductRepository.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return product, nil
}
