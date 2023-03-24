package model

import (
	"gorm.io/gorm"
	domain2 "transfeera.backend.developer.test/src/api/v1/domain"
)

type Beneficiaries []Beneficiary

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

func (b Beneficiary) ToDomain() domain2.Beneficiary {
	return domain2.Beneficiary{
		Status:         b.Status,
		Name:           b.Name,
		DocumentNumber: b.DocumentNumber,
		Email:          b.Email,
		PixKeyType:     b.PixKeyType,
		PixKeyValue:    b.PixKeyValue,
		BankInfo: domain2.BankInfo{
			Bank:    b.Bank,
			Agency:  b.Agency,
			Account: b.Account,
		},
	}
}

func (Beneficiary) FromDomain(b domain2.Beneficiary) Beneficiary {
	return Beneficiary{
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

func (b Beneficiaries) ToDomain() []domain2.Beneficiary {
	beneficiaries := make([]domain2.Beneficiary, 0)
	for _, beneficiary := range b {
		beneficiaries = append(beneficiaries, domain2.Beneficiary{
			Status:         beneficiary.Status,
			Name:           beneficiary.Name,
			DocumentNumber: beneficiary.DocumentNumber,
			Email:          beneficiary.Email,
			PixKeyType:     beneficiary.PixKeyType,
			PixKeyValue:    beneficiary.PixKeyValue,
			BankInfo: domain2.BankInfo{
				Bank:    beneficiary.Bank,
				Agency:  beneficiary.Agency,
				Account: beneficiary.Account,
			},
		})
	}
	return beneficiaries
}
