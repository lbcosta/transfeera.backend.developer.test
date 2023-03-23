package provider

import (
	"context"
	"go.uber.org/fx"
	"transfeera.backend.developer.test/api/handlers"
	repositories "transfeera.backend.developer.test/api/repositories/adapters"
	"transfeera.backend.developer.test/api/services"
)

type AppOptions struct {
	Port   string
	Router interface{}
}

type Api interface {
	Start(port string) error
}

func NewApp(options AppOptions) *fx.App {
	appPort := options.Port
	return fx.New(
		fx.Provide(fx.Annotate(options.Router, fx.As(new(Api)))),
		fx.Provide(providers()...),
		fx.Invoke(func(lifecycle fx.Lifecycle, api Api) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(context.Context) error {
						go api.Start(appPort)
						return nil
					},
				},
			)
		}),
		fx.NopLogger,
	)
}

func providers() []interface{} {
	return []interface{}{
		handlers.NewCreateBeneficiaryHandler,
		handlers.NewGetBeneficiariesHandler,
		handlers.NewUpdateBeneficiaryHandler,
		handlers.NewDeleteBeneficiariesHandler,
		services.NewGetBeneficiariesService,
		repositories.NewBeneficiaryRepository,
	}
}
