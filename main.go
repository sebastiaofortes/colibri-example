package main

import (
	"github.com/colibri-project-io/colibri-sdk-go/pkg/di"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

func start(r []restserver.Route) int32 {
	restserver.AddRoutes(r)
	restserver.ListenAndServe()

	return 0
}

func main() {
	container := di.NewContainer()
	container.AddGlobalDependencies([]interface{}{})
	container.StartApp(start)
}
