package producthandler_test

import (
	"go-api/cmd/database"
	"go-api/handler/producthandler"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	db, err := database.InitDatabase()
	if err != nil {
		log.Panic("Failed to connect to the database:", err)
	}

	producthandler := producthandler.NewProductHandler(db)

	app := fiber.New()

	t.Run("TestGetAllProducts", func(t *testing.T) {
		app.Get("/products", producthandler.GetAll)

		req := httptest.NewRequest("GET", "/products", nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("TestGetByID", func(t *testing.T) {
		app.Get("/products/:id", producthandler.GetByID)

		req := httptest.NewRequest("GET", "/products/1", nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("TestGetByIDErrorNotFound", func(t *testing.T) {
		app.Get("/products/:id", producthandler.GetByID)

		req := httptest.NewRequest("GET", "/products/999", nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	t.Run("TestGetByIDErrorBadRequest", func(t *testing.T) {
		app.Get("/products/:id", producthandler.GetByID)

		req := httptest.NewRequest("GET", "/products/abc", nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
