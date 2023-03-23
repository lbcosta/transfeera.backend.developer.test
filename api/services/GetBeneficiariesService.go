package services

import (
	"transfeera.backend.developer.test/api/domain"
	"transfeera.backend.developer.test/api/repositories"
)

type GetBeneficiariesService struct {
	beneficiaryRepository repositories.BeneficiaryRepository
}

func NewGetBeneficiariesService(beneficiaryRepository repositories.BeneficiaryRepository) GetBeneficiariesService {
	return GetBeneficiariesService{beneficiaryRepository: beneficiaryRepository}
}

func (s GetBeneficiariesService) Call(filter string) ([]domain.Beneficiary, error) {
	beneficiaries, err := s.beneficiaryRepository.Get(filter)
	if err != nil {
		return nil, err
	}
	return beneficiaries, nil
}
