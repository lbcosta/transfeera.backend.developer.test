package main

import (
	"transfeera.backend.developer.test/src/api"
	"transfeera.backend.developer.test/src/provider"
)

func main() {
	app := provider.NewApp(provider.AppOptions{
		Port:   ":8080",
		Router: api.NewRouter,
	})

	app.Run()
}
