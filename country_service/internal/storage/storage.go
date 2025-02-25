package storage

import (
	"context"
	"country_service/internal/domain/models"
	postgresql "country_service/internal/storage/postgreSQL"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*postgresql.Queries
	log *logrus.Logger
	pool *pgxpool.Pool
}

// Repository constructor
func NewRepository(
	pgxpool *pgxpool.Pool,
	log *logrus.Logger,
) Repository {
	return &repo{
		Queries: postgresql.New(pgxpool),
		log:     log,
		pool:	 pgxpool,
	}
}

// TODO test sql statements
// Func for work with DB
type Repository interface {
	GetCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error)
	GetAllCountry(ctx context.Context) (countries []*models.Country, err error)
	CreateCountry(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error)
	UpdateCountrybyID(ctx context.Context, country *models.Country) (err error)
	DeleteCountrybyID(ctx context.Context, country_id int) (err error)
}

// Пока пользуемся принципом PDD
// CreateCountry implements Repository.
func (r *repo) CreateCountry(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error) {
	sqlStatement := `INSERT INTO country (country_title, country_capital, country_area) VALUES ($1, $2, $3) RETURNING country_id`

	country_id := 0
	err = r.pool.QueryRow(ctx, sqlStatement, country_title, country_capital, country_area).Scan(&country_id)
	if err != nil {
		return nil, fmt.Errorf("can`t create country: %w", err)
	}

	country, err = r.GetCountrybyID(ctx, country_id)
	if err != nil {
		return nil, fmt.Errorf("can`t find country: %w", err)
	}

	return country, nil 
}

// DeleteCountrybyID implements Repository.
func (r *repo) DeleteCountrybyID(ctx context.Context, country_id int) (err error) {
	sqlStatement := `DELETE FROM country WHERE country_id=$1`
	
	_, err = r.pool.Exec(ctx, sqlStatement, country_id)
	if err != nil {
		return fmt.Errorf("can`t delete country: %w", err)
	}

	return err
}

// GetAllCountry implements Repository.
func (r *repo) GetAllCountry(ctx context.Context) (countries []*models.Country, err error) {
	sqlStatement := `SELECT * FROM country`
	
	rows, err := r.pool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("can`t query country list: %w", err)
	}

	for rows.Next(){
		country := &models.Country{}
		err := rows.Scan(
			&country.Country_id,
			&country.Country_title,
			&country.Country_capital,
			&country.Country_area,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		countries = append(countries, country)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return countries, err
}

// GetCountrybyID implements Repository.
func (r *repo) GetCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error) {
	sqlStatement := `SELECT * FROM country WHERE country_id=$1`

	country = &models.Country{}
	err = r.pool.QueryRow(ctx, sqlStatement, country_id).Scan(
		&country.Country_id,
		&country.Country_title,
		&country.Country_capital,
		&country.Country_area,
	)
	if err != nil {
		return nil, fmt.Errorf("Couldn`t find country: %w", err)
	}

	return country, nil
}

// UpdateCountrybyID implements Repository.
func (r *repo) UpdateCountrybyID(ctx context.Context, country *models.Country) (err error) {
	sqlStatement := `UPDATE country SET country_title=$2, country_capital=$3, country_area=$4 WHERE country_id=$1`

	_, err = r.pool.Exec(ctx, sqlStatement, country.Country_id, country.Country_title, country.Country_capital, country.Country_area)
	if err != nil {
		return fmt.Errorf("can`t update country: %w", err)
	}

	return nil 
}

func NewStorage(ctx context.Context, dsn string, log *logrus.Logger) (Repository, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return NewRepository(pool, log), nil
}
