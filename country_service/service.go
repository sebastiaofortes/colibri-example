package country_service

type CountryService struct{}

func NewCountryService() CountryService {
	return CountryService{}
}

func (c *CountryService) GetAll() string{
	return "response countries"
}