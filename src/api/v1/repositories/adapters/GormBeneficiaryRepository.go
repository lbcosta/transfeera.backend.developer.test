package adapters

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/repositories"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
	"transfeera.backend.developer.test/src/api/v1/repositories/queries"
	"transfeera.backend.developer.test/src/config"
)

const ResourceNotFoundErr = "resource not found"

type GormBeneficiaryRepository struct {
	databaseConnection config.Database
}

func NewBeneficiaryRepository(databaseConnection config.Database) repositories.BeneficiaryRepository {
	return GormBeneficiaryRepository{databaseConnection: databaseConnection}
}

func (r GormBeneficiaryRepository) GetByID(id uint) (*model.Beneficiary, error) {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return nil, err
	}
	defer r.databaseConnection.Disconnect(database)

	var beneficiary model.Beneficiary
	err = database.First(&beneficiary, id).Error
	if err != nil {
		return nil, err
	}

	return &beneficiary, nil
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

func (r GormBeneficiaryRepository) Update(data *model.Beneficiary) (*model.Beneficiary, error) {
	database, err := r.databaseConnection.Connect()
	if err != nil {
		return nil, err
	}
	defer r.databaseConnection.Disconnect(database)

	err = database.Save(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
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
