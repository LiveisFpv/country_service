package app

import (
	"country_service/internal/services/country"
	postgresql "country_service/internal/storage/postgreSQL"
)

type App struct {
	gRPCServer *country.Server
	storage    postgresql.Repository
}

func NewApp(storage postgresql.Repository, gRPCServer *country.Server) {
	return &App{
		gRPCServer: gRPCServer,
		storage:    storage,
	}
}
