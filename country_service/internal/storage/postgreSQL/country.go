package postgresql

import (
	"context"
	"country_service/internal/domain/models"
	"fmt"
)

func (r *Queries) CreateCountry(ctx context.Context, country_title, country_capital, country_area string) (country *models.Country, err error) {
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
func (r *Queries) DeleteCountrybyID(ctx context.Context, country_id int) (err error) {
	sqlStatement := `DELETE FROM country WHERE country_id=$1`

	_, err = r.pool.Exec(ctx, sqlStatement, country_id)
	if err != nil {
		return fmt.Errorf("can`t delete country: %w", err)
	}

	return err
}

// GetAllCountry implements Repository.
func (r *Queries) GetAllCountry(ctx context.Context, pagination *models.Pagination, filters []*models.Filter, orderbies []*models.OrderBy) (countries []*models.Country, paginate *models.Pagination, err error) {
	sqlStatement := `SELECT * FROM country`

	//TODO filter, paginate, orderby before
	rows, err := r.pool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, nil, fmt.Errorf("can`t query country list: %w", err)
	}

	for rows.Next() {
		country := &models.Country{}
		err := rows.Scan(
			&country.Country_id,
			&country.Country_title,
			&country.Country_capital,
			&country.Country_area,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("can`t process query result: %w", err)
		}
		countries = append(countries, country)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return countries, paginate, err
}

// GetCountrybyID implements Repository.
func (r *Queries) GetCountrybyID(ctx context.Context, country_id int) (country *models.Country, err error) {
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
func (r *Queries) UpdateCountrybyID(ctx context.Context, country *models.Country) (err error) {
	sqlStatement := `UPDATE country SET country_title=$2, country_capital=$3, country_area=$4 WHERE country_id=$1`

	_, err = r.pool.Exec(ctx, sqlStatement, country.Country_id, country.Country_title, country.Country_capital, country.Country_area)
	if err != nil {
		return fmt.Errorf("can`t update country: %w", err)
	}

	return nil
}
