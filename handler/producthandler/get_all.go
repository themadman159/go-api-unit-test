package producthandler

import "github.com/gofiber/fiber/v2"

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		return h.Response.InternalServer(c, err.Error())
	}
	return h.Response.Success(c, "", products)
}
