package city_service

type CityService struct{}

func NewCityService() CityService {
	return CityService{}
}

func (c *CityService) GetAll() string{
	return "response cities"
}