package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	PixKeyValidator         = "validate_pix_key_value"
	DocumentNumberValidator = "validate_document_number"
)

func ValidateDocumentNumber(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$|^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`).MatchString(fl.Field().String())
}

func ErrorMessage(validationErrors validator.ValidationErrors) error {
	errorMsg := "error on the following fields: "

	for idx, err := range validationErrors {
		errorMsg += err.Field()

		if idx != len(validationErrors)-1 {
			errorMsg += ", "
		}
	}

	return errors.New(errorMsg)
}
