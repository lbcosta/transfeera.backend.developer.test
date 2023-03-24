package services

import (
	"errors"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/repositories"
)

const InvalidPixErr = "pix key type and pix key value do not match"

type UpdateBeneficiaryService struct {
	beneficiaryRepository repositories.BeneficiaryRepository
}

func NewUpdateBeneficiaryService(beneficiaryRepository repositories.BeneficiaryRepository) UpdateBeneficiaryService {
	return UpdateBeneficiaryService{beneficiaryRepository: beneficiaryRepository}
}

func (s UpdateBeneficiaryService) Call(id int, data request.UpdateBeneficiaryRequest) (*domain.Beneficiary, error) {
	beneficiary, err := s.beneficiaryRepository.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	beneficiary.Email = updateNonEmpty(beneficiary.Email, data.Email)

	if beneficiary.Status == domain.StatusRascunho {
		beneficiary.Name = updateNonEmpty(beneficiary.Name, data.Name)
		beneficiary.DocumentNumber = updateNonEmpty(beneficiary.DocumentNumber, data.DocumentNumber)
		beneficiary.PixKeyType = updateNonEmpty(beneficiary.PixKeyType, data.PixKeyType)
		beneficiary.PixKeyValue = updateNonEmpty(beneficiary.PixKeyValue, data.PixKeyValue)

		if data.IsPixUpdated() && !beneficiary.IsPixValid() {
			return nil, errors.New(InvalidPixErr)
		}
	}

	updatedBeneficiary, err := s.beneficiaryRepository.Update(beneficiary)
	if err != nil {
		return nil, err
	}

	domainUpdatedBeneficiary := updatedBeneficiary.ToDomain()

	return &domainUpdatedBeneficiary, nil
}

func updateNonEmpty(prev, newValue string) string {
	if newValue == "" {
		return prev
	}
	return newValue
}
