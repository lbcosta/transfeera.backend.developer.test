package services

import (
	"transfeera.backend.developer.test/src/api/repositories"
)

type DeleteBeneficiariesService struct {
	beneficiaryRepository repositories.BeneficiaryRepository
}

func NewDeleteBeneficiariesService(beneficiaryRepository repositories.BeneficiaryRepository) DeleteBeneficiariesService {
	return DeleteBeneficiariesService{beneficiaryRepository: beneficiaryRepository}
}

func (s DeleteBeneficiariesService) Call(ids []uint) error {
	err := s.beneficiaryRepository.Delete(ids)
	if err != nil {
		return err
	}
	return nil
}
