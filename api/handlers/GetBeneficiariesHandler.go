package handlers

import "github.com/gofiber/fiber/v2"

type GetBeneficiariesHandler struct {
}

func NewGetBeneficiariesHandler() GetBeneficiariesHandler {
	return GetBeneficiariesHandler{}
}

// Handle
// query by status, name, pixKeyType or pixKeyValue
// paginate by 10
// metadata
func (h GetBeneficiariesHandler) Handle(c *fiber.Ctx) error {
	return c.SendString("get beneficiaries")
}
