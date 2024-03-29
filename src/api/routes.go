package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"transfeera.backend.developer.test/src/api/v1/handlers"
)

type Router struct {
	CreateBeneficiary   handlers.CreateBeneficiaryHandler
	DeleteBeneficiaries handlers.DeleteBeneficiariesHandler
	GetBeneficiaries    handlers.GetBeneficiariesHandler
	UpdateBeneficiary   handlers.UpdateBeneficiaryHandler
}

func NewRouter(createBeneficiary handlers.CreateBeneficiaryHandler, deleteBeneficiaries handlers.DeleteBeneficiariesHandler, getBeneficiaries handlers.GetBeneficiariesHandler, updateBeneficiary handlers.UpdateBeneficiaryHandler) Router {
	return Router{CreateBeneficiary: createBeneficiary, DeleteBeneficiaries: deleteBeneficiaries, GetBeneficiaries: getBeneficiaries, UpdateBeneficiary: updateBeneficiary}
}

func (r Router) Start(port string) error {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	v1 := app.Group("api/v1/beneficiaries")
	v1.Get("/", r.GetBeneficiaries.Handle)
	v1.Post("/", r.CreateBeneficiary.Handle)
	v1.Patch("/:id", r.UpdateBeneficiary.Handle)
	v1.Delete("/", r.DeleteBeneficiaries.Handle)

	return app.Listen(port)
}
