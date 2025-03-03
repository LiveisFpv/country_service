package app

import (
	"context"
	"country_service/internal/grpc/countrygrpc"
	"country_service/internal/services/country"
	"country_service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	GRPCServer *countrygrpc.App
	Storage    storage.Repository
}

// Constructor APP creates gRPCServer, storage
func New(
	ctx context.Context,
	log *logrus.Logger,
	grpcPort int,
	dsn string,
	tokenTTL time.Duration,
) *App {
	storage, err := storage.NewStorage(ctx, dsn, log)
	if err != nil {
		panic(err)
	}

	//Todo service
	countryService := country.New(log, storage, tokenTTL)
	grpcApp := countrygrpc.New(log, countryService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
		Storage:    storage,
	}
}
