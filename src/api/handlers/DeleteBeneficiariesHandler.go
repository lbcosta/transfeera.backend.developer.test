package handlers

import (
	"github.com/gofiber/fiber/v2"
	"transfeera.backend.developer.test/src/api/handlers/request"
	"transfeera.backend.developer.test/src/api/handlers/response"
	"transfeera.backend.developer.test/src/api/services"
)

type DeleteBeneficiariesHandler struct {
	deleteBeneficiaries services.DeleteBeneficiariesService
}

func NewDeleteBeneficiariesHandler(deleteBeneficiaries services.DeleteBeneficiariesService) DeleteBeneficiariesHandler {
	return DeleteBeneficiariesHandler{deleteBeneficiaries: deleteBeneficiaries}
}

func (h DeleteBeneficiariesHandler) Handle(c *fiber.Ctx) error {
	var req request.DeleteBeneficiariesRequest

	if err := c.BodyParser(&req); err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	err := h.deleteBeneficiaries.Call(req.Ids)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
