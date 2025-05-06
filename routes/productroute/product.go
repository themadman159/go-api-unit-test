package productroute

import (
	"go-api/handler/producthandler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRoute(api fiber.Router, db *gorm.DB) {

	productHandler := producthandler.NewProductHandler(db)
	api.Group("/products").
		Get("/", productHandler.GetAll)
}
