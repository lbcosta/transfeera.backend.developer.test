package repositories

import (
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/model"
)

type BeneficiaryRepository interface {
	Get(filter string) (model.Beneficiaries, error)
	Create(data domain.Beneficiary) (*model.Beneficiary, error)
	Update(data domain.Beneficiary) (*model.Beneficiary, error)
	Delete(ids []uint) error
}
