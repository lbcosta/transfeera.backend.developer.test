package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"transfeera.backend.developer.test/src/api/v1/handlers/response"
	repositories "transfeera.backend.developer.test/src/api/v1/repositories/adapters"
	"transfeera.backend.developer.test/src/api/v1/services"
	"transfeera.backend.developer.test/src/config"
	mocks "transfeera.backend.developer.test/src/mocks/src/api/v1/repositories"
)

type GetBeneficiariesTestSuite struct {
	suite.Suite
	SomeError             error
	App                   *fiber.App
	beneficiaryRepository *mocks.BeneficiaryRepository
}

func (suite *GetBeneficiariesTestSuite) SetupTest() {
	mockedBeneficiaryRepository := new(mocks.BeneficiaryRepository)

	getBeneficiariesHandler := NewGetBeneficiariesHandler(services.NewGetBeneficiariesService(repositories.NewBeneficiaryRepository(config.NewPostgresDatabase())))
	mockedGetBeneficiariesHandler := NewGetBeneficiariesHandler(services.NewGetBeneficiariesService(mockedBeneficiaryRepository))

	app := fiber.New()
	app.Get("/", getBeneficiariesHandler.Handle)
	app.Get("/mock", mockedGetBeneficiariesHandler.Handle)

	suite.App = app
	suite.beneficiaryRepository = mockedBeneficiaryRepository
	suite.SomeError = errors.New("some error")
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_Success() {
	req := httptest.NewRequest("GET", "/", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var getBeneficiariesResponse response.GetBeneficiariesResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &getBeneficiariesResponse)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalCount, 37)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalPages, 4)
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_SearchByStatus() {
	req := httptest.NewRequest("GET", "/?filter=Validado", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var getBeneficiariesResponse response.GetBeneficiariesResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &getBeneficiariesResponse)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), len(getBeneficiariesResponse.Data), 10)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalCount, 20)
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_SearchByName() {
	req := httptest.NewRequest("GET", "/?filter=Alex%20Johnson", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var getBeneficiariesResponse response.GetBeneficiariesResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &getBeneficiariesResponse)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalCount, 1)
	assert.Equal(suite.T(), getBeneficiariesResponse.Data[0].DocumentNumber, "258147369")
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_SearchByPixKeyType() {
	req := httptest.NewRequest("GET", "/?filter=EMAIL", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var getBeneficiariesResponse response.GetBeneficiariesResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &getBeneficiariesResponse)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), len(getBeneficiariesResponse.Data), 9)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalCount, 9)
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_SearchByPixKeyValue() {
	req := httptest.NewRequest("GET", "/?filter=bobsmith@example.com", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var getBeneficiariesResponse response.GetBeneficiariesResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &getBeneficiariesResponse)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), getBeneficiariesResponse.Metadata.TotalCount, 1)
	assert.Equal(suite.T(), getBeneficiariesResponse.Data[0].DocumentNumber, "456789123")
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_NonExistingPage() {
	req := httptest.NewRequest("GET", "/?page=5", nil)

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

	assert.Equal(suite.T(), fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(suite.T(), errResp.Status, response.StatusInvalidInput)
	assert.Equal(suite.T(), errResp.Error, "The requested page does not exist.")
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_ServiceFails() {
	suite.beneficiaryRepository.On("Get", mock.Anything).Return(nil, suite.SomeError)

	req := httptest.NewRequest("GET", "/mock", nil)

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
	assert.Equal(suite.T(), errResp.Status, response.StatusError)
	assert.Equal(suite.T(), errResp.Error, "some error")
}

func TestGetBeneficiariesTestSuite(t *testing.T) {
	suite.Run(t, new(GetBeneficiariesTestSuite))
}
