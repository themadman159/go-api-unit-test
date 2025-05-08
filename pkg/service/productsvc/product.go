package productsvc

import (
	"go-api/pkg/model"
	"go-api/pkg/repository/productrepo"

	"gorm.io/gorm"
)

type IProductService interface {
	GetAllProducts() (*[]model.Product, error)
	GetByID(id int) (*model.Product, error)
}

type ProductService struct {
	ProductRepository productrepo.IProductRepository
}

func NewProductService(dbconn *gorm.DB) IProductService {
	return &ProductService{
		ProductRepository: productrepo.NewProductRepository(dbconn),
	}
}
