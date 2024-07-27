package city

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CityController struct {
}

func NewCityController() *CityController {
	return &CityController{}
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
	wctx.JsonResponse(http.StatusOK, "OK")
}