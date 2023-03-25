package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
	mocks "transfeera.backend.developer.test/src/mocks/src/api/v1/repositories"
)

type UpdateBeneficiaryTestSuite struct {
	suite.Suite
	SomeError                error
	beneficiaryRepository    *mocks.BeneficiaryRepository
	updateBeneficiaryService UpdateBeneficiaryService
	beneficiaryModel         model.Beneficiary
	beneficiaryDomain        domain.Beneficiary
	bankInfo                 domain.BankInfo
	updateBeneficiaryRequest request.UpdateBeneficiaryRequest
}

func (suite *UpdateBeneficiaryTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.beneficiaryRepository = new(mocks.BeneficiaryRepository)
	suite.updateBeneficiaryService = NewUpdateBeneficiaryService(suite.beneficiaryRepository)
	suite.beneficiaryModel = model.Beneficiary{
		Status:         "Rascunho",
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
		Bank:           "TransfeeraBank",
		Agency:         "0123-4",
		Account:        "987654-3",
	}
	suite.bankInfo = domain.BankInfo{
		Bank:    "TransfeeraBank",
		Agency:  "0123-4",
		Account: "987654-3",
	}
	suite.beneficiaryDomain = domain.Beneficiary{
		Status:         "Rascunho",
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
		BankInfo:       suite.bankInfo,
	}
	suite.updateBeneficiaryRequest = request.UpdateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
	}
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_Success() {
	id := 1
	suite.beneficiaryRepository.On("GetByID", uint(id)).Return(&suite.beneficiaryModel, nil)
	suite.beneficiaryRepository.On("Update", &suite.beneficiaryModel).Return(&suite.beneficiaryModel, nil)

	beneficiary, err := suite.updateBeneficiaryService.Call(id, suite.updateBeneficiaryRequest)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &suite.beneficiaryDomain, beneficiary)
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_RequestHasEmptyValues() {
	id := 1
	suite.beneficiaryRepository.On("GetByID", uint(id)).Return(&suite.beneficiaryModel, nil)
	suite.beneficiaryRepository.On("Update", &suite.beneficiaryModel).Return(&suite.beneficiaryModel, nil)

	suite.updateBeneficiaryRequest.Name = ""
	beneficiary, err := suite.updateBeneficiaryService.Call(id, suite.updateBeneficiaryRequest)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &suite.beneficiaryDomain, beneficiary)
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_Error() {
	id := 1
	suite.beneficiaryRepository.On("GetByID", uint(id)).Return(&suite.beneficiaryModel, nil)
	suite.beneficiaryRepository.On("Update", &suite.beneficiaryModel).Return(nil, suite.SomeError)

	_, err := suite.updateBeneficiaryService.Call(id, suite.updateBeneficiaryRequest)

	assert.Error(suite.T(), err)
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_BeneficiaryDoesNotExists() {
	id := 1
	suite.beneficiaryRepository.On("GetByID", uint(id)).Return(nil, suite.SomeError)

	_, err := suite.updateBeneficiaryService.Call(id, suite.updateBeneficiaryRequest)

	assert.Error(suite.T(), err)
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_PixTypeAndPixValueDoesNotMatch() {
	id := 1
	suite.beneficiaryRepository.On("GetByID", uint(id)).Return(&suite.beneficiaryModel, nil)

	suite.updateBeneficiaryRequest.PixKeyValue = "85999999999"
	_, err := suite.updateBeneficiaryService.Call(id, suite.updateBeneficiaryRequest)

	assert.Error(suite.T(), err)
}

func TestUpdateBeneficiaryTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateBeneficiaryTestSuite))
}
