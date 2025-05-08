package producthandler

import (
	"go-api/pkg/service/productsvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IProductHandler interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
}

type ProductHandler struct {
	ProductService productsvc.IProductService
	Response       responseutil.IResponseUtil
}

func NewProductHandler(dbconn *gorm.DB) IProductHandler {
	return &ProductHandler{
		ProductService: productsvc.NewProductService(dbconn),
		Response:       responseutil.NewResponseUtil(),
	}
}
