package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"transfeera.backend.developer.test/src/api/v1/domain"
	"transfeera.backend.developer.test/src/api/v1/handlers/request"
	"transfeera.backend.developer.test/src/api/v1/handlers/response"
	repositories "transfeera.backend.developer.test/src/api/v1/repositories/adapters"
	"transfeera.backend.developer.test/src/api/v1/services"
	"transfeera.backend.developer.test/src/config"
	mocks "transfeera.backend.developer.test/src/mocks/src/api/v1/repositories"
)

type CreateBeneficiaryTestSuite struct {
	suite.Suite
	SomeError                   error
	App                         *fiber.App
	mockedBeneficiaryRepository *mocks.BeneficiaryRepository
	domainBeneficiary           *domain.Beneficiary
}

func (suite *CreateBeneficiaryTestSuite) SetupTest() {
	mockedBeneficiaryRepository := new(mocks.BeneficiaryRepository)
	mockedCreateBeneficiariesHandler := NewCreateBeneficiaryHandler(services.NewCreateBeneficiaryService(mockedBeneficiaryRepository), services.NewGetBankInfoService())

	createBeneficiaryHandler := NewCreateBeneficiaryHandler(services.NewCreateBeneficiaryService(repositories.NewBeneficiaryRepository(config.NewTestDatabase())), services.NewGetBankInfoService())

	app := fiber.New()
	app.Post("/", createBeneficiaryHandler.Handle)
	app.Post("/mock", mockedCreateBeneficiariesHandler.Handle)

	suite.App = app
	suite.mockedBeneficiaryRepository = mockedBeneficiaryRepository
	suite.SomeError = errors.New("some error")
	suite.domainBeneficiary = &domain.Beneficiary{
		Status:         "Rascunho",
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
		BankInfo: domain.BankInfo{
			Bank:    "TransfeeraBank",
			Agency:  "1234-5",
			Account: "987654-3",
		},
	}
}

func (suite *CreateBeneficiaryTestSuite) TearDownTest() {
	config.Destroy()
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiaries_Success() {
	reqBody := request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var createdBeneficiary domain.Beneficiary
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &createdBeneficiary)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), suite.domainBeneficiary, &createdBeneficiary)
	assert.Equal(suite.T(), domain.StatusRascunho, createdBeneficiary.Status)
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiaries_InvalidPixType() {
	reqBody := request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "INVALIDO",
		PixKeyValue:    "lbcosta.dev@gmail.com",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqBodyJSON))
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

	assert.Equal(suite.T(), fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(suite.T(), response.StatusInvalidInput, errResp.Status)
	assert.Equal(suite.T(), "error on the following fields: PixKeyType, PixKeyValue", errResp.Error)
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiaries_InvalidPixValue() {
	reqBody := request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "04788380340",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqBodyJSON))
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

	assert.Equal(suite.T(), fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(suite.T(), response.StatusInvalidInput, errResp.Status)
	assert.Equal(suite.T(), "error on the following fields: PixKeyValue", errResp.Error)
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiaries_InvalidEmail() {
	reqBody := request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev",
		PixKeyType:     "CPF",
		PixKeyValue:    "04788380340",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqBodyJSON))
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

	assert.Equal(suite.T(), fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(suite.T(), response.StatusInvalidInput, errResp.Status)
	assert.Equal(suite.T(), "error on the following fields: Email", errResp.Error)
}

func (suite *CreateBeneficiaryTestSuite) TestCreateBeneficiaries_ServiceFails() {
	suite.mockedBeneficiaryRepository.On("Create", *suite.domainBeneficiary).Return(nil, suite.SomeError)

	reqBody := request.CreateBeneficiaryRequest{
		Name:           "Leonardo Costa",
		DocumentNumber: "04788380340",
		Email:          "lbcosta.dev@gmail.com",
		PixKeyType:     "EMAIL",
		PixKeyValue:    "lbcosta.dev@gmail.com",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("POST", "/mock", bytes.NewBuffer(reqBodyJSON))
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
	assert.Equal(suite.T(), "some error", errResp.Error)
}

func TestCreateBeneficiaryTestSuite(t *testing.T) {
	suite.Run(t, new(CreateBeneficiaryTestSuite))
}
