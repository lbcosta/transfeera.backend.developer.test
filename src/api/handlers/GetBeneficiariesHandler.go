package handlers

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
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
		resError := response.GetBeneficiariesError{
			Status: response.StatusError,
			Code:   http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.Status(http.StatusUnprocessableEntity).JSON(resError)
	}

	if len(beneficiaries.Data) == 0 {
		resError := response.GetBeneficiariesError{
			Status: response.StatusInvalidInput,
			Code:   http.StatusBadRequest,
			Error:  "The requested page does not exist.",
		}
		return c.Status(http.StatusBadRequest).JSON(resError)
	}

	res := response.GetBeneficiariesResponse{
		Status:   response.StatusSuccess,
		Code:     http.StatusOK,
		Metadata: response.NewMetadata(beneficiaries.TotalCount, page, PerPage),
		Data:     beneficiaries.Data,
	}

	return c.JSON(res)
}
