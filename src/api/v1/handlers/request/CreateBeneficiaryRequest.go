package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type CreateBeneficiaryRequest struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required,validate_document_number"`
	Email          string `json:"email" validate:"required,email,max=250"`
	PixKeyType     string `json:"pix_key_type" validate:"required,oneof=CPF CNPJ EMAIL TELEFONE CHAVE_ALEATORIA"`
	PixKeyValue    string `json:"pix_key_value" validate:"required,max=140,validate_pix_key_value"`
}

func ValidatePixKeyValue(fl validator.FieldLevel) bool {
	request := fl.Parent().Interface().(CreateBeneficiaryRequest)
	switch request.PixKeyType {
	case "CPF":
		return regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`).MatchString(fl.Field().String())
	case "CNPJ":
		return regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`).MatchString(fl.Field().String())
	case "EMAIL":
		return regexp.MustCompile(`^[a-z0-9+_.-]+@[a-z0-9.-]+$`).MatchString(fl.Field().String())
	case "TELEFONE":
		return regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`).MatchString(fl.Field().String())
	case "CHAVE_ALEATORIA":
		return regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`).MatchString(fl.Field().String())
	default:
		return false
	}
}

func ValidateDocumentNumber(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$|^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`).MatchString(fl.Field().String())
}
