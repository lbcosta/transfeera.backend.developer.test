package model

import (
	"gorm.io/gorm"
	"transfeera.backend.developer.test/src/api/domain"
)

type Beneficiary struct {
	gorm.Model
	Status         string `json:"status"`
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
	Email          string `json:"email"`
	PixKeyType     string `json:"pix_key_type"`
	PixKeyValue    string `json:"pix_key_value"`
	Bank           string `json:"bank"`
	Agency         string `json:"agency"`
	Account        string `json:"account"`
}

func (b Beneficiary) ToDomain() domain.Beneficiary {
	return domain.Beneficiary{
		Status:         b.Status,
		Name:           b.Name,
		DocumentNumber: b.DocumentNumber,
		Email:          b.Email,
		PixKeyType:     b.PixKeyType,
		PixKeyValue:    b.PixKeyValue,
		Bank:           b.Bank,
		Agency:         b.Agency,
		Account:        b.Account,
	}
}

type Beneficiaries []Beneficiary

func (b Beneficiaries) ToDomain() []domain.Beneficiary {
	var beneficiaries []domain.Beneficiary
	for _, beneficiary := range b {
		beneficiaries = append(beneficiaries, domain.Beneficiary{
			Status:         beneficiary.Status,
			Name:           beneficiary.Name,
			DocumentNumber: beneficiary.DocumentNumber,
			Email:          beneficiary.Email,
			PixKeyType:     beneficiary.PixKeyType,
			PixKeyValue:    beneficiary.PixKeyValue,
			Bank:           beneficiary.Bank,
			Agency:         beneficiary.Agency,
			Account:        beneficiary.Account,
		})
	}
	return beneficiaries
}