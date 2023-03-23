package services

import (
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/repositories"
)

type BeneficiariesData struct {
	Data       []domain.Beneficiary `json:"data"`
	TotalCount int                  `json:"total_count"`
}

type GetBeneficiariesService struct {
	beneficiaryRepository repositories.BeneficiaryRepository
}

func NewGetBeneficiariesService(beneficiaryRepository repositories.BeneficiaryRepository) GetBeneficiariesService {
	return GetBeneficiariesService{beneficiaryRepository: beneficiaryRepository}
}

func (s GetBeneficiariesService) Call(filter string, page, perPage int) (*BeneficiariesData, error) {
	beneficiaries, err := s.beneficiaryRepository.Get(filter)
	if err != nil {
		return nil, err
	}

	beneficiariesPage := paginate(beneficiaries.ToDomain(), (page-1)*perPage, perPage)

	data := &BeneficiariesData{
		Data:       beneficiariesPage,
		TotalCount: len(beneficiaries),
	}

	return data, nil
}

func paginate(x []domain.Beneficiary, offset int, limit int) []domain.Beneficiary {
	if offset > len(x) {
		offset = len(x)
	}

	end := offset + limit
	if end > len(x) {
		end = len(x)
	}

	return x[offset:end]
}
