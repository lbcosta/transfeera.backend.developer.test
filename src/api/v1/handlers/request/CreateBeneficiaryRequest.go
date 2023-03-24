package request

import (
	"github.com/go-playground/validator/v10"
	"transfeera.backend.developer.test/src/util"
)

type CreateBeneficiaryRequest struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required,validate_document_number"`
	Email          string `json:"email" validate:"required,email,max=250"`
	PixKeyType     string `json:"pix_key_type" validate:"required,oneof=CPF CNPJ EMAIL TELEFONE CHAVE_ALEATORIA"`
	PixKeyValue    string `json:"pix_key_value" validate:"required,max=140,validate_pix_key_value"`
}

func (r CreateBeneficiaryRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation(PixKeyValidator, validatePixKeyValue)
	validate.RegisterValidation(DocumentNumberValidator, ValidateDocumentNumber)

	err := validate.Struct(r)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ErrorMessage(validationErrors)
	}

	return nil
}

func validatePixKeyValue(fl validator.FieldLevel) bool {
	request := fl.Parent().Interface().(CreateBeneficiaryRequest)
	return util.ValidatePix(request.PixKeyType, fl.Field().String())
}
