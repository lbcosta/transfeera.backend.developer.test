package repositories

import (
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
)

type BeneficiaryRepository interface {
	GetByID(id uint) (*model.Beneficiary, error)
	Get(filter string) (model.Beneficiaries, error)
	Create(data domain.Beneficiary) (*model.Beneficiary, error)
	Update(data *model.Beneficiary) (*model.Beneficiary, error)
	Delete(ids []uint) error
}
