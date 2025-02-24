package grpc

import (
	country_v1 "github.com/LiveisFpv/country_v1/gen/go/country"
)

type serverAPI struct {
	country_v1.UnimplementedCountryServer
	country Country
}

type Country interface {
}
