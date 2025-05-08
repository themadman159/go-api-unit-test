package routes

import (
	"go-api/routes/example"
	"go-api/routes/productroute"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	example.ExampleRoute(api, db)
	productroute.ProductRoute(api, db)
}
