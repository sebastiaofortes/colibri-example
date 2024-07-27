package main

import (
	"os"

	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/di"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/sebastiaofortes/colibri-example/city_controller"
	"github.com/sebastiaofortes/colibri-example/country_controller"
	"github.com/sebastiaofortes/colibri-example/region_controller"
	"github.com/sebastiaofortes/colibri-example/city_service"
	"github.com/sebastiaofortes/colibri-example/country_service"
	"github.com/sebastiaofortes/colibri-example/region_service"
)

func init() {
	os.Setenv("ENVIRONMENT", "development")
	os.Setenv("APP_NAME", "development")
	os.Setenv("APP_TYPE", "serverless")
	os.Setenv("CLOUD", "aws")
	colibri.InitializeApp()
}

func start(routes ...[]restserver.Route) int32 {
	for _, r := range routes {
		restserver.AddRoutes(r)
	}

	restserver.ListenAndServe()

	return 0
}

func main() {
	dependencies := []interface{}{
		// routes
		country_controller.Routes,		
		city_controller.Routes,
		region_controller.Routes,
		// controllers
		region_controller.NewRegionController,
		country_controller.NewCountryController,
		city_controller.NewCityController,
		// services
		city_service.NewCityService,
		country_service.NewCountryService,
		region_service.NewRegionService,
	}
	
	container := di.NewContainer()
	container.AddGlobalDependencies(dependencies)
	container.StartApp(start)
}
