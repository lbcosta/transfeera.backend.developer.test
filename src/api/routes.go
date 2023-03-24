package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	handlers2 "transfeera.backend.developer.test/src/api/v1/handlers"
)

type Router struct {
	CreateBeneficiary   handlers2.CreateBeneficiaryHandler
	DeleteBeneficiaries handlers2.DeleteBeneficiariesHandler
	GetBeneficiaries    handlers2.GetBeneficiariesHandler
	UpdateBeneficiary   handlers2.UpdateBeneficiaryHandler
}

func NewRouter(createBeneficiary handlers2.CreateBeneficiaryHandler, deleteBeneficiaries handlers2.DeleteBeneficiariesHandler, getBeneficiaries handlers2.GetBeneficiariesHandler, updateBeneficiary handlers2.UpdateBeneficiaryHandler) Router {
	return Router{CreateBeneficiary: createBeneficiary, DeleteBeneficiaries: deleteBeneficiaries, GetBeneficiaries: getBeneficiaries, UpdateBeneficiary: updateBeneficiary}
}

func (r Router) Start(port string) error {
	app := fiber.New()

	app.Use(recover.New())

	v1 := app.Group("api/v1/beneficiaries")
	v1.Get("/", r.GetBeneficiaries.Handle)
	v1.Post("/", r.CreateBeneficiary.Handle)
	v1.Patch("/:id", r.UpdateBeneficiary.Handle)
	v1.Delete("/", r.DeleteBeneficiaries.Handle)

	return app.Listen(port)
}
