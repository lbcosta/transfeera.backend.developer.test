package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/handlers/response"
	"transfeera.backend.developer.test/src/api/v1/services"
)

const (
	PixKeyValidator         = "validate_pix_key_value"
	DocumentNumberValidator = "validate_document_number"
)

type CreateBeneficiaryHandler struct {
	createBeneficiary services.CreateBeneficiaryService
	getBankInfo       services.GetBankInfoService
}

func NewCreateBeneficiaryHandler(createBeneficiary services.CreateBeneficiaryService, getBankInfo services.GetBankInfoService) CreateBeneficiaryHandler {
	return CreateBeneficiaryHandler{createBeneficiary: createBeneficiary, getBankInfo: getBankInfo}
}

func (h CreateBeneficiaryHandler) Handle(c *fiber.Ctx) error {
	var req request.CreateBeneficiaryRequest

	if err := c.BodyParser(&req); err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	validate := validator.New()
	validate.RegisterValidation(PixKeyValidator, request.ValidatePixKeyValue)
	validate.RegisterValidation(DocumentNumberValidator, request.ValidateDocumentNumber)

	err := validate.Struct(req)
	if err != nil {
		errorResponse := response.ErrorResponse{}.FromValidation(err.(validator.ValidationErrors))
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	bankInfo, err := h.getBankInfo.Call(req.PixKeyValue)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Status: response.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorResponse)
	}

	beneficiary, err := h.createBeneficiary.Call(req, bankInfo)
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
