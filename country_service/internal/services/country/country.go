package country

import (
	"context"
	"country_service/internal/domain/models"
	"time"

	"github.com/sirupsen/logrus"
)

//TODO translate requset from handlers to db

// All methods
type CountryStorage interface {
	GetCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error)
	GetAllCountry(ctx context.Context) (countries []*models.Country, err error)
	CreateCountry(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error)
	UpdateCountrybyID(ctx context.Context, country *models.Country) (err error)
	DeleteCountrybyID(ctx context.Context, country_id int) (err error)
}

type CountryService struct {
	log            *logrus.Logger
	countryStorage CountryStorage
	tokenTTL       time.Duration
}

// Constructor service of Country
func New(
	log *logrus.Logger,
	countryStorage CountryStorage,
	tokenTTL time.Duration,
) *CountryService {
	return &CountryService{
		log:            log,
		countryStorage: countryStorage,
		tokenTTL:       tokenTTL,
	}
}

// Add_Country implements countrygrpc.Country.
func (c *CountryService) Add_Country(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error) {
	panic("unimplemented")
}

// Delete_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Delete_CountrybyID(ctx context.Context, country_id int) (err error) {
	panic("unimplemented")
}

// Get_All_Country implements countrygrpc.Country.
func (c *CountryService) Get_All_Country(ctx context.Context) (countries []*models.Country, err error) {
	panic("unimplemented")
}

// Get_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Get_CountrybyID(ctx context.Context, country_id int) (country *models.Country, err error) {
	panic("unimplemented")
}

// Update_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Update_CountrybyID(ctx context.Context, country *models.Country) (err error) {
	panic("unimplemented")
}
