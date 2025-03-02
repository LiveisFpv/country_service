package main

import (
	"context"
	"country_service/internal/app"
	"country_service/internal/config"
	"country_service/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

// TODO start microservice
func main() {
	//TODO init config obj
	cfg := config.MustLoad()

	//TODO init logger
	//if need more opt make create in lib/logger
	log := logger.LoggerSetup(true)

	//TODO init app
	ctx := context.Background()
	app := app.New(ctx, log, cfg.GRPC.Port, cfg.Dsn, cfg.TokenTTL)
	log.Info("Start service")
	//TODO start grpc-Server
	go func() {
		app.GRPCServer.MustRun()
	}()

	//TODO graceful shutdown
	// Graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	// Waiting for SIGINT (pkill -2) or SIGTERM
	<-stop
	// initiate graceful shutdown
	app.GRPCServer.Stop()
	log.Info("GRPCserver stopped")
	app.Storage.Stop()
	log.Info("Postgres connection closed")
}
