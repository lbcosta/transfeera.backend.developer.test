package request

import (
	"github.com/go-playground/validator/v10"
)

type UpdateBeneficiaryRequest struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number,omitempty" validate:"omitempty,validate_document_number"`
	Email          string `json:"email,omitempty" validate:"omitempty,email,max=250"`
	PixKeyType     string `json:"pix_key_type,omitempty" validate:"omitempty,oneof=CPF CNPJ EMAIL TELEFONE CHAVE_ALEATORIA"`
	PixKeyValue    string `json:"pix_key_value,omitempty" validate:"omitempty,max=140"`
}

func (r UpdateBeneficiaryRequest) IsPixUpdated() bool {
	return r.PixKeyType != "" || r.PixKeyValue != ""
}

func (r UpdateBeneficiaryRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation(DocumentNumberValidator, ValidateDocumentNumber)

	err := validate.Struct(r)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ErrorMessage(validationErrors)
	}

	return nil
}
