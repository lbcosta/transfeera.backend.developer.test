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

type UpdateBeneficiaryTestSuite struct {
	suite.Suite
	SomeError                   error
	App                         *fiber.App
	mockedBeneficiaryRepository *mocks.BeneficiaryRepository
	domainBeneficiary           *domain.Beneficiary
}

func (suite *UpdateBeneficiaryTestSuite) SetupTest() {
	mockedBeneficiaryRepository := new(mocks.BeneficiaryRepository)

	updateBeneficiaryHandler := NewUpdateBeneficiaryHandler(services.NewUpdateBeneficiaryService(repositories.NewBeneficiaryRepository(config.NewTestDatabase())))
	mockedUpdateBeneficiariesHandler := NewUpdateBeneficiaryHandler(services.NewUpdateBeneficiaryService(mockedBeneficiaryRepository))

	app := fiber.New()
	app.Patch("/:id", updateBeneficiaryHandler.Handle)
	app.Patch("/mock/:id", mockedUpdateBeneficiariesHandler.Handle)

	suite.App = app
	suite.mockedBeneficiaryRepository = mockedBeneficiaryRepository
	suite.SomeError = errors.New("some error")
	suite.domainBeneficiary = &domain.Beneficiary{
		Status:         "Validado",
		Name:           "John Doe",
		DocumentNumber: "12345678907",
		Email:          "johndoe@example.com",
		PixKeyType:     "CPF",
		PixKeyValue:    "12345678907",
		BankInfo: domain.BankInfo{
			Bank:    "ABC Bank",
			Agency:  "1234",
			Account: "56789",
		},
	}
}

func (suite *UpdateBeneficiaryTestSuite) TearDownTest() {
	config.Destroy()
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_Success() {
	const Id = "1"
	const email = "johndoe@gmail.com"
	reqBody := request.UpdateBeneficiaryRequest{
		Email: email,
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("PATCH", "/"+Id, bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var updatedBeneficiary domain.Beneficiary
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &updatedBeneficiary)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.domainBeneficiary.Email = email

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), suite.domainBeneficiary, &updatedBeneficiary)
}

func (suite *UpdateBeneficiaryTestSuite) TestUpdateBeneficiary_StatusValidadoCanOnlyEditEmail() {
	const Id = "1"
	reqBody := request.UpdateBeneficiaryRequest{
		PixKeyValue: "12345678902",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	req := httptest.NewRequest("PATCH", "/"+Id, bytes.NewBuffer(reqBodyJSON))
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
	assert.Equal(suite.T(), services.ShouldNotUpdateErr, errResp.Error)
}

func TestUpdateBeneficiaryTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateBeneficiaryTestSuite))
}
