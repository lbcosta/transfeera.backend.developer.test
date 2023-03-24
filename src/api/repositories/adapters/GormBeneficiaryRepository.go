package adapters

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"transfeera.backend.developer.test/src/api/domain"
	"transfeera.backend.developer.test/src/api/repositories"
	"transfeera.backend.developer.test/src/api/repositories/model"
	"transfeera.backend.developer.test/src/api/repositories/queries"
	"transfeera.backend.developer.test/src/config"
)

const ResourceNotFoundErr = "resource not found"

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

func (r GormBeneficiaryRepository) Create(data domain.Beneficiary) (*model.Beneficiary, error) {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return nil, err
	}
	defer r.databaseConnection.Disconnect(database)

	beneficiary := model.Beneficiary{}.FromDomain(data)

	err = database.Create(&beneficiary).Error
	if err != nil {
		return nil, err
	}

	return &beneficiary, nil
}

func (r GormBeneficiaryRepository) Update(data domain.Beneficiary) (*model.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (r GormBeneficiaryRepository) Delete(ids []uint) error {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return err
	}
	defer r.databaseConnection.Disconnect(database)

	err = database.Transaction(func(tx *gorm.DB) error {
		var beneficiaries []model.Beneficiary
		result := tx.Find(&beneficiaries, ids)
		if result.RowsAffected == 0 {
			return errors.New(ResourceNotFoundErr)
		}

		err = tx.Delete(&model.Beneficiary{}, ids).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
