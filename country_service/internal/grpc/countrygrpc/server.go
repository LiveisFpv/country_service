package countrygrpc

import (
	country_v1 "github.com/LiveisFpv/country_v1/gen/go/country"
	"google.golang.org/grpc"
)

type serverAPI struct {
	country_v1.UnimplementedCountryServer
	country Country
}

type Country interface {
}

// It how constructor but not constructor:В
func Register(gRPCServer *grpc.Server, country Country) {
	country_v1.RegisterCountryServer(gRPCServer, &serverAPI{country: country})
}
