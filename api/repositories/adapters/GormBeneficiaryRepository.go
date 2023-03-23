package adapters

import (
	"transfeera.backend.developer.test/api/domain"
	"transfeera.backend.developer.test/api/repositories"
	"transfeera.backend.developer.test/config"
)

type GormBeneficiaryRepository struct {
	databaseConnection config.PostgresDatabase
}

func NewBeneficiaryRepository(databaseConnection config.PostgresDatabase) repositories.BeneficiaryRepository {
	return GormBeneficiaryRepository{databaseConnection: databaseConnection}
}

func (r GormBeneficiaryRepository) Get(filter string) ([]domain.Beneficiary, error) {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return nil, err
	}
	defer r.databaseConnection.Disconnect(database)

	var beneficiaries []domain.Beneficiary
	err = database.Find(&beneficiaries).Error
	if err != nil {
		return nil, err
	}

	return beneficiaries, nil
}

func (r GormBeneficiaryRepository) Create(beneficiary domain.Beneficiary) (domain.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (r GormBeneficiaryRepository) Update(beneficiary domain.Beneficiary) (domain.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (r GormBeneficiaryRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
