package adapters

import (
	"strings"
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/model"
	"transfeera.backend.developer.test/src/api/repositories"
	"transfeera.backend.developer.test/src/api/repositories/queries"
	"transfeera.backend.developer.test/src/config"
)

type GormBeneficiaryRepository struct {
	databaseConnection config.PostgresDatabase
}

func NewBeneficiaryRepository(databaseConnection config.PostgresDatabase) repositories.BeneficiaryRepository {
	return GormBeneficiaryRepository{databaseConnection: databaseConnection}
}

func (r GormBeneficiaryRepository) Get(filter string) (model.Beneficiaries, error) {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return nil, err
	}
	defer r.databaseConnection.Disconnect(database)

	var beneficiaries []model.Beneficiary
	err = database.Raw(strings.Replace(queries.Filter, "${SEARCH}", filter, 1)).Find(&beneficiaries).Error
	if err != nil {
		return nil, err
	}

	return beneficiaries, nil
}

func (r GormBeneficiaryRepository) Create(beneficiary domain.Beneficiary) (model.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (r GormBeneficiaryRepository) Update(beneficiary domain.Beneficiary) (model.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (r GormBeneficiaryRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
