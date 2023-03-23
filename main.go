package main

import (
	"transfeera.backend.developer.test/api"
	"transfeera.backend.developer.test/provider"
)

func main() {
	app := provider.NewApp(provider.AppOptions{
		Port:   ":3000",
		Router: api.NewRouter,
	})

	app.Run()
}
