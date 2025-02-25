package storage

import (
	"country_service/internal/domain/models"
	postgresql "country_service/internal/storage/postgreSQL"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*postgresql.Queries
	log *logrus.Logger
}

// Repository constructor
func NewRepository(
	pgxpool *pgxpool.Pool,
	log *logrus.Logger,
) Repository {
	return &repo{
		Queries: New(pgxpool),
		log:     log,
	}
}

// TODO sql statement and func
// Func for work with DB
type Repository interface {
	GetCountrybyID(country_id int) (country *models.Country, err error)
	GetAllCountry() (countries *[]models.Country, err error)
	CreateCountry(country *models.Country) (country_id int, err error)
	UpdateCountrybyID(country *models.Country) (err error)
	DeleteCountrybyID(country_id int) (err error)
}
