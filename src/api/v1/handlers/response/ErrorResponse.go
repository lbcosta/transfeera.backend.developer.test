package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

func (r ErrorResponse) FromValidation(validationErrors validator.ValidationErrors) ErrorResponse {
	errorMsg := "error on the following fields: "
	for idx, err := range validationErrors {
		errorMsg += err.Field()

		if idx != len(validationErrors)-1 {
			errorMsg += ", "
		}
	}

	return ErrorResponse{
		Status: StatusInvalidInput,
		Code:   fiber.StatusBadRequest,
		Error:  errorMsg,
	}
}
