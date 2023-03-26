package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/handlers/response"
	repositories "transfeera.backend.developer.test/src/api/v1/repositories/adapters"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
	"transfeera.backend.developer.test/src/api/v1/services"
	"transfeera.backend.developer.test/src/config"
)

type DeleteBeneficiaryTestSuite struct {
	suite.Suite
	SomeError         error
	App               *fiber.App
	TestDatabase      config.Database
	domainBeneficiary *domain.Beneficiary
}

func (suite *DeleteBeneficiaryTestSuite) SetupTest() {
	testDatabase := config.NewTestDatabase()
	deleteBeneficiaryHandler := NewDeleteBeneficiariesHandler(services.NewDeleteBeneficiariesService(repositories.NewBeneficiaryRepository(testDatabase)))

	app := fiber.New()
	app.Delete("/", deleteBeneficiaryHandler.Handle)

	suite.App = app
	suite.TestDatabase = testDatabase
	suite.SomeError = errors.New("some error")
}

func (suite *DeleteBeneficiaryTestSuite) TearDownTest() {
	config.Destroy()
}

func (suite *DeleteBeneficiaryTestSuite) TestDeleteBeneficiaries_Success() {
	const Id = 1

	reqBody := request.DeleteBeneficiariesRequest{
		Ids: []uint{Id},
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("DELETE", "/", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	db, err := suite.TestDatabase.Connect()
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}
	defer suite.TestDatabase.Disconnect(db)

	var beneficiary model.Beneficiary
	err = db.First(&beneficiary, Id).Error

	assert.Equal(suite.T(), fiber.StatusNoContent, resp.StatusCode)
	assert.Error(suite.T(), err)
	assert.True(suite.T(), errors.Is(err, gorm.ErrRecordNotFound))
}

func (suite *DeleteBeneficiaryTestSuite) TestDeleteBeneficiaries_NotFound() {
	reqBody := request.DeleteBeneficiariesRequest{
		Ids: []uint{1131, 1132, 1133},
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("DELETE", "/", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var errResp response.ErrorResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(suite.T(), response.StatusError, errResp.Status)
	assert.Equal(suite.T(), "resource not found", errResp.Error)
}

func TestDeleteBeneficiaryTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteBeneficiaryTestSuite))
}
