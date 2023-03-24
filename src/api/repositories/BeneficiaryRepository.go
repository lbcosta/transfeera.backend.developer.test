package repositories

import (
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/model"
)

type BeneficiaryRepository interface {
	Get(filter string) (model.Beneficiaries, error)
	Create(beneficiary domain.Beneficiary) (model.Beneficiary, error)
	Update(beneficiary domain.Beneficiary) (model.Beneficiary, error)
	Delete(ids []uint) error
}
