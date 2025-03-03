package countrygrpc

import (
	"context"
	"country_service/internal/domain/models"
	"fmt"

	country_v1 "github.com/LiveisFpv/country_v1/gen/go/country"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get_CountryById(ctx context.Context, req *country_v1.Get_CountryById_Request) (*country_v1.Get_CountryById_Response, error) {
	if req.CountryId < 1 {
		return nil, status.Error(codes.InvalidArgument, "country_id is required")
	}

	//Send request to Service
	country, err := s.country.Get_CountrybyID(ctx, int(req.CountryId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}
	return &country_v1.Get_CountryById_Response{
		CountryId:      int64(country.Country_id),
		CountryTitle:   country.Country_title,
		CountryCapital: country.Country_capital,
		CountryArea:    country.Country_area,
	}, nil
}

// TODO Pag,Sort,Filter
func (s *serverAPI) Get_All_Country(
	ctx context.Context,
	req *country_v1.Get_All_Country_Request,
) (*country_v1.Get_All_Country_Response, error) {
	pagination := &models.Pagination{
		Current: int(req.Pagination.Current),
		Limit:   int(req.Pagination.Limit),
		Total:   int(req.Pagination.Total),
	}
	filters := []*models.Filter{}
	for _, filter := range req.Filters {
		filters = append(filters, &models.Filter{
			Field: filter.Field,
			Value: filter.Value,
		})
	}
	orderbies := []*models.OrderBy{}
	for _, orderby := range req.Orderby {
		orderbies = append(orderbies, &models.OrderBy{
			Field:     orderby.Field,
			Direction: orderby.Direction,
		})
	}
	countries, paginate, err := s.country.Get_All_Country(ctx, pagination, filters, orderbies)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}
	resp := &country_v1.Get_All_Country_Response{}
	for _, country := range countries {
		resp.Countries = append(resp.Countries,
			&country_v1.Get_CountryById_Response{
				CountryId:      int64(country.Country_id),
				CountryTitle:   country.Country_title,
				CountryCapital: country.Country_capital,
				CountryArea:    country.Country_area,
			},
		)
	}
	resp.Pagination.Current = int64(paginate.Current)
	resp.Pagination.Limit = int64(paginate.Limit)
	resp.Pagination.Total = int64(paginate.Total)
	return resp, nil
}

func (s *serverAPI) Add_Country(
	ctx context.Context,
	req *country_v1.Add_Country_Request,
) (*country_v1.Add_Country_Response, error) {

	if len(req.CountryTitle) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Title is required")
	}

	if len(req.CountryCapital) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Capitral is required")
	}

	if len(req.CountryArea) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Area is required")
	}

	country, err := s.country.Add_Country(ctx, req.CountryTitle, req.CountryCapital, req.CountryArea)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &country_v1.Add_Country_Response{
		CountryId: int64(country.Country_id),
	}, nil

}

func (s *serverAPI) Update_CountryById(
	ctx context.Context,
	req *country_v1.Update_CountryById_Request,
) (*country_v1.Update_CountryById_Response, error) {
	if req.CountryId < 1 {
		return nil, status.Error(codes.InvalidArgument, "country_id is required")
	}
	if len(req.CountryTitle) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Title is required")
	}

	if len(req.CountryCapital) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Capitral is required")
	}

	if len(req.CountryArea) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Country Area is required")
	}

	country := &models.Country{
		Country_id:      int(req.CountryId),
		Country_title:   req.CountryTitle,
		Country_capital: req.CountryCapital,
		Country_area:    req.CountryArea,
	}
	err := s.country.Update_CountrybyID(ctx, country)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	// Я кнчн не эксперт, но перекидывать данные с запроса на возврат немного странно
	// Повторно дергаем по Id, возвращаем что реально получилось в БД
	updated_country, err := s.country.Get_CountrybyID(ctx, int(req.CountryId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}
	return &country_v1.Update_CountryById_Response{
		CountryTitle:   updated_country.Country_title,
		CountryCapital: updated_country.Country_capital,
		CountryArea:    updated_country.Country_area,
	}, nil
}

func (s *serverAPI) Delete_CountryById(
	ctx context.Context,
	req *country_v1.Delete_CountryById_Request,
) (*country_v1.Delete_CountryById_Response, error) {
	if req.CountryId < 1 {
		return nil, status.Error(codes.InvalidArgument, "country_id is required")
	}

	// Скозали вернуть название, ну я и верну
	// а query лучше лишний раз не трогать
	country, err := s.country.Get_CountrybyID(ctx, int(req.CountryId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}

	err = s.country.Delete_CountrybyID(ctx, int(req.CountryId))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &country_v1.Delete_CountryById_Response{
		CountryTitle:   country.Country_title,
		CountryCapital: country.Country_capital,
		CountryArea:    country.Country_area,
	}, nil
}
