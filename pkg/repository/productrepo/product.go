package productrepo

import (
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAllProducts() (*[]model.Product, error)
	GetByID(id int) (*model.Product, error)
}

type ProductRepository struct {
	Database *gorm.DB
}

func NewProductRepository(dbconn *gorm.DB) IProductRepository {
	return &ProductRepository{
		Database: dbconn,
	}
}
