package util

import "regexp"

func ValidatePix(pixKeyType, pixKeyValue string) bool {
	switch pixKeyType {
	case "CPF":
		return regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`).MatchString(pixKeyValue)
	case "CNPJ":
		return regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`).MatchString(pixKeyValue)
	case "EMAIL":
		return regexp.MustCompile(`^[a-z0-9+_.-]+@[a-z0-9.-]+$`).MatchString(pixKeyValue)
	case "TELEFONE":
		return regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`).MatchString(pixKeyValue)
	case "CHAVE_ALEATORIA":
		return regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`).MatchString(pixKeyValue)
	default:
		return false
	}
}
