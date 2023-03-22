package handlers

import "github.com/gofiber/fiber/v2"

type CreateBeneficiaryHandler struct {
}

func NewCreateBeneficiaryHandler() CreateBeneficiaryHandler {
	return CreateBeneficiaryHandler{}
}

// Handle
// body: name/razãoSocial, cpf/cnpj, email, pixKeyType, pixKeyValue (, bank, agency, account?)
// validation middleware
// create with status: Rascunho
func (h CreateBeneficiaryHandler) Handle(c *fiber.Ctx) error {
	return c.SendString("create beneficiary")
}
