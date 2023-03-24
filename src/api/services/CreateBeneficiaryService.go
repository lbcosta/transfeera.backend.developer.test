package services

import (
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/handlers/request"
	"transfeera.backend.developer.test/src/api/repositories"
)

type CreateBeneficiaryService struct {
	beneficiaryRepository repositories.BeneficiaryRepository
}

func NewCreateBeneficiaryService(beneficiaryRepository repositories.BeneficiaryRepository) CreateBeneficiaryService {
	return CreateBeneficiaryService{beneficiaryRepository: beneficiaryRepository}
}

func (s CreateBeneficiaryService) Call(data request.CreateBeneficiaryRequest, bankInfo domain.BankInfo) (*domain.Beneficiary, error) {
	beneficiaryData := domain.Beneficiary{
		Status:         domain.StatusRascunho,
		Name:           data.Name,
		DocumentNumber: data.DocumentNumber,
		Email:          data.Email,
		PixKeyType:     data.PixKeyType,
		PixKeyValue:    data.PixKeyValue,
		BankInfo:       bankInfo,
	}

	beneficiary, err := s.beneficiaryRepository.Create(beneficiaryData)
	if err != nil {
		return nil, err
	}

	domainBeneficiary := beneficiary.ToDomain()

	return &domainBeneficiary, nil
}
