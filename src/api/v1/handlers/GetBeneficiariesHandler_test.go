package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
	repositories "transfeera.backend.developer.test/src/api/v1/repositories/adapters"
	"transfeera.backend.developer.test/src/api/v1/services"
	"transfeera.backend.developer.test/src/config"
)

type GetBeneficiariesTestSuite struct {
	suite.Suite
	App              *fiber.App
	PostgresDatabase config.PostgresDatabase
}

func (suite *GetBeneficiariesTestSuite) SetupTest() {
	postgresDatabase := config.NewPostgresDatabase()
	beneficiaryRepository := repositories.NewBeneficiaryRepository(postgresDatabase)
	getBeneficiariesService := services.NewGetBeneficiariesService(beneficiaryRepository)
	getBeneficiariesHandler := NewGetBeneficiariesHandler(getBeneficiariesService)

	app := fiber.New()
	app.Get("/", getBeneficiariesHandler.Handle)

	suite.App = app
	suite.PostgresDatabase = postgresDatabase
}

func (suite *GetBeneficiariesTestSuite) TestGetBeneficiaries_Success() {
	req := httptest.NewRequest("GET", "/", nil)

	resp, err := suite.App.Test(req, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

func TestGetBeneficiariesTestSuite(t *testing.T) {
	suite.Run(t, new(GetBeneficiariesTestSuite))
}
