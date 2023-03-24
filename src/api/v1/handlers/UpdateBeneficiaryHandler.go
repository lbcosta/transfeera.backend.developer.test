package handlers

import (
	"github.com/gofiber/fiber/v2"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/handlers/response"
	"transfeera.backend.developer.test/src/api/v1/services"
)

type UpdateBeneficiaryHandler struct {
	updateBeneficiary services.UpdateBeneficiaryService
}

func NewUpdateBeneficiaryHandler(updateBeneficiary services.UpdateBeneficiaryService) UpdateBeneficiaryHandler {
	return UpdateBeneficiaryHandler{updateBeneficiary: updateBeneficiary}
}

func (h UpdateBeneficiaryHandler) Handle(c *fiber.Ctx) error {
	beneficiaryId, err := c.ParamsInt("id")
	if err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	var req request.UpdateBeneficiaryRequest

	if err = c.BodyParser(&req); err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if err = req.Validate(); err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	beneficiary, err := h.updateBeneficiary.Call(beneficiaryId, req)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse)
	}

	return c.JSON(beneficiary)
}
