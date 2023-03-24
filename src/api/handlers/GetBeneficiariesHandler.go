package handlers

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"transfeera.backend.developer.test/src/api/handlers/response"
	"transfeera.backend.developer.test/src/api/services"
)

const PerPage = 10

type GetBeneficiariesHandler struct {
	getBeneficiaries services.GetBeneficiariesService
}

func NewGetBeneficiariesHandler(getBeneficiaries services.GetBeneficiariesService) GetBeneficiariesHandler {
	return GetBeneficiariesHandler{getBeneficiaries: getBeneficiaries}
}

func (h GetBeneficiariesHandler) Handle(c *fiber.Ctx) error {
	filter := c.Query("filter")
	page := c.QueryInt("page", 1)
	page = int(math.Max(float64(page), 1))

	beneficiaries, err := h.getBeneficiaries.Call(filter, page, PerPage)
	if err != nil {
		resError := response.ErrorResponse{
			Status: response.StatusError,
			Code:   fiber.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resError)
	}

	metadata := response.NewMetadata(beneficiaries.TotalCount, page, PerPage)

	if beneficiaries.TotalCount > 0 && page > metadata.TotalPages {
		errorResponse := response.ErrorResponse{
			Status: response.StatusInvalidInput,
			Code:   fiber.StatusBadRequest,
			Error:  "The requested page does not exist.",
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	res := response.GetBeneficiariesResponse{
		Status:   response.StatusSuccess,
		Code:     fiber.StatusOK,
		Metadata: metadata,
		Data:     beneficiaries.Data,
	}

	return c.JSON(res)
}
