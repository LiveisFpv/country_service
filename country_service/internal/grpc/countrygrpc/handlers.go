package countrygrpc

import (
	"context"
	"fmt"

	country_v1 "github.com/LiveisFpv/country_v1/gen/go/country"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetCountrybyID(
	ctx context.Context,
	req country_v1.Get_CountryById_Requset,
) (*country_v1.Get_CountryById_Response, error) {
	if req.CountryId < 1 {
		return nil, status.Error(codes.InvalidArgument, "country_id is required")
	}

	//Send request to Service
	country, err := s.country.Get_CountrybyID(ctx, int(req.CountryId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}
	return &country_v1.Get_CountryById_Response{
		CountryTitle:   country.Country_title,
		CountryCapital: country.Country_capital,
		CountryArea:    country.Country_area,
	}, nil
}
