package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	mocks "transfeera.backend.developer.test/src/mocks/src/api/v1/repositories"
)

type DeleteBeneficiariesTestSuite struct {
	suite.Suite
	SomeError                  error
	beneficiaryRepository      *mocks.BeneficiaryRepository
	deleteBeneficiariesService DeleteBeneficiariesService
}

func (suite *DeleteBeneficiariesTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.beneficiaryRepository = new(mocks.BeneficiaryRepository)
	suite.deleteBeneficiariesService = NewDeleteBeneficiariesService(suite.beneficiaryRepository)
}

func (suite *DeleteBeneficiariesTestSuite) TestDeleteBeneficiaries_Success() {
	ids := []uint{1, 2, 3}
	suite.beneficiaryRepository.On("Delete", ids).Return(nil)

	err := suite.deleteBeneficiariesService.Call(ids)

	assert.NoError(suite.T(), err)
}

func (suite *DeleteBeneficiariesTestSuite) TestDeleteBeneficiaries_Error() {
	ids := []uint{1, 2, 3}
	suite.beneficiaryRepository.On("Delete", ids).Return(suite.SomeError)

	err := suite.deleteBeneficiariesService.Call(ids)

	assert.Error(suite.T(), err)
}

func TestDeleteBeneficiariesTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteBeneficiariesTestSuite))
}
