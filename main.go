package main

import (
	"transfeera.backend.developer.test/api"
	"transfeera.backend.developer.test/api/handlers"
	"transfeera.backend.developer.test/provider"
)

func main() {
	app := provider.NewApp(provider.AppOptions{
		Port:      ":3000",
		Router:    api.NewRouter,
		Providers: providers(),
	})

	app.Run()
}

func providers() []interface{} {
	return []interface{}{
		handlers.NewCreateBeneficiaryHandler,
		handlers.NewGetBeneficiariesHandler,
		handlers.NewUpdateBeneficiaryHandler,
		handlers.NewDeleteBeneficiariesHandler,
	}
}
