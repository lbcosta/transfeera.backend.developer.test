package handlers

import "github.com/gofiber/fiber/v2"

type UpdateBeneficiaryHandler struct {
}

func NewUpdateBeneficiaryHandler() UpdateBeneficiaryHandler {
	return UpdateBeneficiaryHandler{}
}

// Handle
// Status==Rascunho -> edição de qlq dado
// Status==Validado -> edição somente do email
// Status não pode ser editado
func (h UpdateBeneficiaryHandler) Handle(c *fiber.Ctx) error {
	return c.SendString("update beneficiary")
}
