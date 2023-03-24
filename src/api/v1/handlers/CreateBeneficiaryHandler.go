package handlers

import (
	"github.com/gofiber/fiber/v2"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	response2 "transfeera.backend.developer.test/src/api/v1/handlers/response"
	services2 "transfeera.backend.developer.test/src/api/v1/services"
)

type CreateBeneficiaryHandler struct {
	createBeneficiary services2.CreateBeneficiaryService
	getBankInfo       services2.GetBankInfoService
}

func NewCreateBeneficiaryHandler(createBeneficiary services2.CreateBeneficiaryService, getBankInfo services2.GetBankInfoService) CreateBeneficiaryHandler {
	return CreateBeneficiaryHandler{createBeneficiary: createBeneficiary, getBankInfo: getBankInfo}
}

func (h CreateBeneficiaryHandler) Handle(c *fiber.Ctx) error {
	var req request.CreateBeneficiaryRequest

	if err := c.BodyParser(&req); err != nil {
		errorResponse := response2.ErrorResponse{
			Status: response2.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	// TODO: Validate

	bankInfo, err := h.getBankInfo.Call(req.PixKeyValue)
	if err != nil {
		errorResponse := response2.ErrorResponse{
			Status: response2.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse)
	}

	beneficiary, err := h.createBeneficiary.Call(req, bankInfo)
	if err != nil {
		errorResponse := response2.ErrorResponse{
			Status: response2.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse)
	}

	return c.JSON(beneficiary)
}
