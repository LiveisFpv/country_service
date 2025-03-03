package country

import (
	"context"
	"country_service/internal/domain/models"
	"time"

	"github.com/sirupsen/logrus"
)

// All methods
type CountryStorage interface {
	GetCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error)
	GetAllCountry(ctx context.Context, pagination *models.Pagination, filter []*models.Filter, orderby []*models.Sort) ([]*models.Country, *models.Pagination, error)
	CreateCountry(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error)
	UpdateCountrybyID(ctx context.Context, country *models.Country) (err error)
	DeleteCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error)
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
	const op = "Country.Create"
	log := c.log.WithFields(
		logrus.Fields{
			"op":      op,
			"title":   country_title,
			"capital": country_capital,
			"area":    country_area,
		},
	)
	log.Info("Start Create Country")

	country, err = c.countryStorage.CreateCountry(ctx, country_title, country_capital, country_area)
	if err != nil {
		c.log.Error("failed to create country", err)
		return nil, err
	}

	return country, nil
}

// Delete_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Delete_CountrybyID(ctx context.Context, country_id int) (*models.Country, error) {
	const op = "Country.Delete"
	log := c.log.WithFields(
		logrus.Fields{
			"op": op,
			"id": country_id,
		},
	)
	log.Info("Start Delete Country")
	res, err := c.countryStorage.DeleteCountrybyID(ctx, country_id)
	if err != nil {
		c.log.Error("failed to delete country", err)
		return nil, err
	}
	return res, nil
}

// Get_All_Country implements countrygrpc.Country.
func (c *CountryService) Get_All_Country(ctx context.Context, pagination *models.Pagination, filter []*models.Filter, orderby []*models.Sort) ([]*models.Country, *models.Pagination, error){
	const op = "Country.GetAll"
	log := c.log.WithFields(
		logrus.Fields{
			"op": op,
		},
	)
	log.Info("Start Get ALL Country")

	countries, new_pagination, err := c.countryStorage.GetAllCountry(ctx, pagination, filter, orderby)
	if err != nil {
		c.log.Error("failed to get all countries", err)
		return nil, nil, err
	}

	return countries, new_pagination, nil
}

// Get_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Get_CountrybyID(ctx context.Context, country_id int) (country *models.Country, err error) {
	const op = "Country.GetbyID"
	log := c.log.WithFields(
		logrus.Fields{
			"op": op,
			"id": country_id,
		},
	)
	log.Info("Start Get by ID Country")
	country, err = c.countryStorage.GetCountrybyID(ctx, country_id)
	if err != nil {
		c.log.Error("failed to get country by id", err)
		return nil, err
	}
	return country, nil
}

// Update_CountrybyID implements countrygrpc.Country.
func (c *CountryService) Update_CountrybyID(ctx context.Context, country *models.Country) (err error) {
	const op = "Country.Update"
	log := c.log.WithFields(
		logrus.Fields{
			"op":      op,
			"id":      country.Country_id,
			"title":   country.Country_title,
			"capital": country.Country_capital,
			"area":    country.Country_area,
		},
	)
	log.Info("Start Update country")
	err = c.countryStorage.UpdateCountrybyID(ctx, country)
	if err != nil {
		c.log.Error("failed to update country", err)
		return err
	}
	return nil
}
