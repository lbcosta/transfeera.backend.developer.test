package services

import (
	"transfeera.backend.developer.test/src/api/v1/domain"
)

type GetBankInfoService struct{}

func NewGetBankInfoService() GetBankInfoService {
	return GetBankInfoService{}
}

func (s GetBankInfoService) Call(pixKeyValue string) (domain.BankInfo, error) {
	return domain.BankInfo{
		Bank:    "TransfeeraBank",
		Agency:  "1234-5",
		Account: "987654-3",
	}, nil
}
