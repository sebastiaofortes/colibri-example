package region_service

type RegionService struct{}

func NewRegionService() RegionService {
	return RegionService{}
}

func (c *RegionService) GetAll() string{
	return "response regions"
}