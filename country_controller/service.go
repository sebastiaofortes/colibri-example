package country_controller

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/sebastiaofortes/colibri-example/country_service"
)

type CountryController struct {
	service country_service.CountryService
}

func NewCountryController(serv country_service.CountryService) *CountryController {
	return &CountryController{
		service: serv,
	}
}

func Routes(c *CountryController) []restserver.Route {
	return []restserver.Route{
		{
			URI:      "countries",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
	}
}

func (c *CountryController) GetAll(wctx restserver.WebContext) {
	response := c.service.GetAll()
	wctx.JsonResponse(http.StatusOK, response)
}