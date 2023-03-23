package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"transfeera.backend.developer.test/api/handlers/response"
	"transfeera.backend.developer.test/api/services"
)

const PerPage = 10

type GetBeneficiariesHandler struct {
	getBeneficiaries services.GetBeneficiariesService
}

func NewGetBeneficiariesHandler(getBeneficiaries services.GetBeneficiariesService) GetBeneficiariesHandler {
	return GetBeneficiariesHandler{getBeneficiaries: getBeneficiaries}
}

// Handle
// query by status, name, pixKeyType or pixKeyValue
// paginate by 10
// metadata
func (h GetBeneficiariesHandler) Handle(c *fiber.Ctx) error {
	filter := c.Query("filter")
	page := c.QueryInt("page", 1)

	beneficiaries, err := h.getBeneficiaries.Call(filter)
	if err != nil {
		resError := response.GetBeneficiariesError{
			Status: response.StatusError,
			Code:   http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}
		return c.JSON(resError)
	}
	
	res := response.GetBeneficiariesResponse{
		Status:   response.StatusSuccess,
		Code:     http.StatusOK,
		Metadata: response.NewMetadata(len(beneficiaries), page, PerPage),
		Data:     beneficiaries,
	}

	return c.JSON(res)
}
