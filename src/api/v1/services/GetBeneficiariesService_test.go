package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
	mocks "transfeera.backend.developer.test/src/mocks/src/api/v1/repositories"
)

type GetBeneficiariesTestSuite struct {
	suite.Suite
	SomeError               error
	beneficiaryRepository   *mocks.BeneficiaryRepository
	getBeneficiariesService GetBeneficiariesService
	beneficiaryModel        model.Beneficiary
	beneficiariesData       BeneficiariesData
}

func (suite *GetBeneficiariesTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.beneficiaryRepository = new(mocks.BeneficiaryRepository)
	suite.getBeneficiariesService = NewGetBeneficiariesService(suite.beneficiaryRepository)
	suite.beneficiaryModel = model.Beneficiary{
		Model:          gorm.Model{},
		Status:         "Validade",
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
		Bank:           "TransfeeraBank",
		Agency:         "0123-4",
		Account:        "987654-3",
	}
	suite.beneficiariesData = BeneficiariesData{
		Data:       []domain.Beneficiary{suite.beneficiaryModel.ToDomain()},
		TotalCount: 1,
	}

}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_Success() {
	page := 1
	perPage := 10
	suite.beneficiaryRepository.On("Get", mock.Anything).Return(model.Beneficiaries{suite.beneficiaryModel}, nil)

	beneficiariesData, err := suite.getBeneficiariesService.Call(mock.Anything, page, perPage)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &suite.beneficiariesData, beneficiariesData)
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_OffsetGreaterThenBeneficiariesLength() {
	page := 2
	perPage := 10
	suite.beneficiaryRepository.On("Get", mock.Anything).Return(model.Beneficiaries{suite.beneficiaryModel}, nil)

	beneficiariesData, err := suite.getBeneficiariesService.Call(mock.Anything, page, perPage)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &BeneficiariesData{Data: []domain.Beneficiary{}, TotalCount: 1}, beneficiariesData)
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_Error() {
	page := 1
	perPage := 10
	suite.beneficiaryRepository.On("Get", mock.Anything).Return(nil, suite.SomeError)

	beneficiariesData, err := suite.getBeneficiariesService.Call(mock.Anything, page, perPage)

	assert.Nil(suite.T(), beneficiariesData)
	assert.Error(suite.T(), err)
}

func TestGetBeneficiariesTestSuite(t *testing.T) {
	suite.Run(t, new(GetBeneficiariesTestSuite))
}
