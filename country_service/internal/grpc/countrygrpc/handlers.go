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
	req *country_v1.Get_CountryById_Requset,
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
func (s *serverAPI) Get_All_Country(
	ctx context.Context,
	req *country_v1.Get_All_Country_Request,
) (*country_v1.Get_All_Country_Response, error) {
	countries, err := s.country.Get_All_Country(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}
	resp := &country_v1.Get_All_Country_Response{}
	for _, country := range countries {
		resp.Countries = append(resp.Countries,
			&country_v1.Get_CountryById_Response{
				CountryTitle:   country.Country_title,
				CountryCapital: country.Country_capital,
				CountryArea:    country.Country_area,
			},
		)
	}
	return resp, nil
}
