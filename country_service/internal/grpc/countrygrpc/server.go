package countrygrpc

import (
	"context"
	"country_service/internal/domain/models"

	country_v1 "github.com/LiveisFpv/country_v1/gen/go/country"
	"google.golang.org/grpc"
)

type serverAPI struct {
	country_v1.UnimplementedCountryServer
	country Country
}

// Methods needed for handlers on Service
type Country interface {
	Get_CountrybyID(
		ctx context.Context,
		country_id int,
	) (country *models.Country, err error)
	Get_All_Country(
		ctx context.Context,
		pagination *models.Pagination,
		filter []*models.Filter,
		orderby []*models.Sort,
	) (
		countries []*models.Country, 
		new_pagination *models.Pagination,
		err error)
	Add_Country(
		ctx context.Context,
		country_title string,
		country_capital string,
		country_area string,
	) (country *models.Country, err error)
	Update_CountrybyID(
		ctx context.Context,
		country *models.Country,
	) (err error)
	Delete_CountrybyID(
		ctx context.Context,
		country_id int,
	) (country *models.Country, err error)
}

// It how constructor but not constructor:В
func Register(gRPCServer *grpc.Server, country Country) {
	country_v1.RegisterCountryServer(gRPCServer, &serverAPI{country: country})
}
