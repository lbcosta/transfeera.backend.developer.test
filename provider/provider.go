package provider

import (
	"context"
	"go.uber.org/fx"
)

type AppOptions struct {
	Port      string
	Router    interface{}
	Providers []interface{}
}

type Api interface {
	Start(port string) error
}

func NewApp(options AppOptions) *fx.App {
	appPort := options.Port
	return fx.New(
		fx.Provide(fx.Annotate(options.Router, fx.As(new(Api)))),
		fx.Provide(options.Providers...),
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
