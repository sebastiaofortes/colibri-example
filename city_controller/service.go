package city_controller

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/sebastiaofortes/colibri-example/city_service"
)

type CityController struct {
	service city_service.CityService
}

func NewCityController(serv city_service.CityService) *CityController {
	return &CityController{
		service: serv,
	}
}

func Routes(c *CityController) []restserver.Route {
	return []restserver.Route{
		{
			URI:      "cities",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
	}
}

func (c *CityController) GetAll(wctx restserver.WebContext) {
	response := c.service.GetAll()
	wctx.JsonResponse(http.StatusOK, response)
}