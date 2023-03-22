package handlers

import "github.com/gofiber/fiber/v2"

type DeleteBeneficiariesHandler struct {
}

func NewDeleteBeneficiariesHandler() DeleteBeneficiariesHandler {
	return DeleteBeneficiariesHandler{}
}

// Handle
// body: array de IDs a serem excluídos
func (h DeleteBeneficiariesHandler) Handle(c *fiber.Ctx) error {
	return c.SendString("delete beneficiaries")
}
