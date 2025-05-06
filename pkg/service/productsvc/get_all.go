package productsvc

import "go-api/pkg/model"

func (s *ProductService) GetAllProducts() (*[]model.Product, error) {
	products, err := s.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
