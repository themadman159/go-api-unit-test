package productrepo_test

import (
	"go-api/cmd/database"
	"go-api/pkg/repository/productrepo"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestProduct(t *testing.T) {
	db, err := database.InitDatabase()
	if err != nil {
		log.Panic("Failed to connect to the database:", err)
	}

	productRepo := productrepo.NewProductRepository(db)

	t.Run("TestGetAllProducts", func(t *testing.T) {
		products, err := productRepo.GetAllProducts()
		if err != nil {
			t.Error("Failed to get products:", err)
		}

		assert.NotNil(t, products)
		assert.Equal(t, 10, len(*products))
	})

	t.Run("TestGetByID", func(t *testing.T) {
		product, err := productRepo.GetByID(1)
		if err != nil {
			t.Error("Failed to get product by ID:", err)
		}

		assert.NotNil(t, product)
		assert.Equal(t, 1, product.ID)
	})

	t.Run("TestGetByIDErrorNotFound", func(t *testing.T) {
		product, err := productRepo.GetByID(0)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		assert.Nil(t, product)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
		assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
	})
}
