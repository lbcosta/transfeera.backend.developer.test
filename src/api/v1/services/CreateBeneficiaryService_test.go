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

type CreateBeneficiaryTestSuite struct {
	suite.Suite
	SomeError                error
	beneficiaryRepository    *mocks.BeneficiaryRepository
	createBeneficiaryService CreateBeneficiaryService
	beneficiaryModel         model.Beneficiary
	beneficiaryDomain        domain.Beneficiary
	bankInfo                 domain.BankInfo
	createBeneficiaryRequest request.CreateBeneficiaryRequest
}

func (suite *CreateBeneficiaryTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.beneficiaryRepository = new(mocks.BeneficiaryRepository)
	suite.createBeneficiaryService = NewCreateBeneficiaryService(suite.beneficiaryRepository)
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
	suite.createBeneficiaryRequest = request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
	}
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiary_Success() {
	suite.beneficiaryRepository.On("Create", suite.beneficiaryDomain).Return(&suite.beneficiaryModel, nil)

	beneficiary, err := suite.createBeneficiaryService.Call(suite.createBeneficiaryRequest, suite.bankInfo)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &suite.beneficiaryDomain, beneficiary)
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiary_Error() {
	suite.beneficiaryRepository.On("Create", suite.beneficiaryDomain).Return(nil, suite.SomeError)

	_, err := suite.createBeneficiaryService.Call(suite.createBeneficiaryRequest, suite.bankInfo)

	assert.Error(suite.T(), err)
}


func TestCreateBeneficiaryTestSuite(t *testing.T) {
	suite.Run(t, new(CreateBeneficiaryTestSuite))
}
