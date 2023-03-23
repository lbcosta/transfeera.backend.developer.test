package repositories

import "transfeera.backend.developer.test/api/domain"

type BeneficiaryRepository interface {
	Get(filter string) ([]domain.Beneficiary, error)
	Create(beneficiary domain.Beneficiary) (domain.Beneficiary, error)
	Update(beneficiary domain.Beneficiary) (domain.Beneficiary, error)
	Delete(id int) error
}
