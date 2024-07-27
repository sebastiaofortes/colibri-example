package region_controller

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/sebastiaofortes/colibri-example/region_service"
)

type RegionController struct {
	service region_service.RegionService
}

func NewRegionController(serv region_service.RegionService) *RegionController {
	return &RegionController{
		service: serv,
	}
}

func Routes(c *RegionController) []restserver.Route {
	return []restserver.Route{
		{
			URI:      "regions",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
	}
}

func (c *RegionController) GetAll(wctx restserver.WebContext) {
	response := c.service.GetAll()
	wctx.JsonResponse(http.StatusOK, response)
}