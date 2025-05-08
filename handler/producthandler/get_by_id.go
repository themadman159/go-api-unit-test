package producthandler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {

	productID, err := c.ParamsInt("id")
	if productID == 0 || err != nil {
		return h.Response.BadRequest(c, "Product ID is required")
	}

	product, err := h.ProductService.GetByID(int(productID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return h.Response.NotFound(c, "Product not found")
		}
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Success(c, "", product)
}
