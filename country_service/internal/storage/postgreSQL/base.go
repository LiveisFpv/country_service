package postgresql

import (
	"context"
	"country_service/internal/domain/models"
	"fmt"
)

func unpackFilter(ctx context.Context, baseQuery string, filters []*models.Filter) (filteredQuery string) {
	for _, filter := range filters {
		baseQuery += fmt.Sprintf(" AND %s = %s", filter.Field, filter.Value)
	}
	return baseQuery
}

func unpackOrder(ctx context.Context, baseQuery string, orderby []*models.Sort) (orderedQuery string) {
	if len(orderby) > 0 {
		// Начинаем сортировку
		baseQuery += " ORDER BY"
		for i, order := range orderby {
			if i > 0 {
				// После первого элемента каждый идет через запятую
				baseQuery += ", "
			}
			baseQuery += fmt.Sprintf("%s %s", order.By, order.Direction)
		}
	}
	return baseQuery
}
